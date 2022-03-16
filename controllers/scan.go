package controllers

import (
	"log"

	"github.com/ShangRui-hash/arp-spoofing-go/logic"
	"github.com/ShangRui-hash/arp-spoofing-go/pkg/privileges"
	"github.com/ShangRui-hash/arp-spoofing-go/pkg/utils"
	"github.com/ShangRui-hash/arp-spoofing-go/settings"

	"github.com/abiosoft/ishell"
)

//ScanHandler 扫描功能处理函数
func ScanHandler(c *ishell.Context) {
	//0.检查用户权限
	if privileges.Check() {
		c.Println("权限不足，不能发包，请先提升权限")
		return
	}
	//1.解析扫描范围
	ipRange, err := settings.Options.Get("range")
	if err != nil {
		log.Println(err)
		return
	}
	ipList, err := utils.GetIPList(ipRange)
	if err != nil {
		log.Printf("utils.GetIPList failed,err:%v\n", err)
		return
	}
	//2.选择扫描方式
	method, err := settings.Options.Get("method")
	if err != nil {
		log.Println(err)
		return
	}
	//3.选择网卡
	ifname, err := settings.Options.Get("ifname")
	if err != nil {
		log.Println(err)
		return
	}

	// 3.业务逻辑层
	if err := logic.Scan(c, ipList, ifname, method); err != nil {
		log.Println(err)
		return
	}
}
