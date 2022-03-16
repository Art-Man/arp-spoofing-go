package settings

import (
	"fmt"

	"github.com/ShangRui-hash/arp-spoofing-go/models"
	"github.com/ShangRui-hash/arp-spoofing-go/pkg/utils"
)

//ARPScanOptions ARP扫描配置
var (
	Options *models.Options
)

//Init 初始化配置
func Init() (err error) {
	//获取默认网卡、扫描范围、网关
	defaultIfname, defaultScanRange, defaultGateway, err := utils.GetDefaultOptions()
	if err != nil {
		fmt.Println("utils.GetDefaultOptions() failed,err:", err)
		return err
	}
	//初始化扫描配置项
	Options = models.NewOptions("ARP Scan Options")
	Options.Add("ifname", defaultIfname, true, "监听哪个网卡")
	Options.Add("range", defaultScanRange, true, "扫描范围")
	Options.Add("method", "arp", true, "扫描方式:all,arp,udp")
	Options.Add("gateway", defaultGateway, true, "局域网的网关")
	return nil
}
