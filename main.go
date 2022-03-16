package main

import (
	"log"

	"github.com/ShangRui-hash/arp-spoofing-go/routers"
	"github.com/ShangRui-hash/arp-spoofing-go/settings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "arp-spoofing-go > ",
	})
	//初始化配置
	if err := settings.Init(); err != nil {
		log.Println("settings.Init failed,err:", err)
		return
	}
	//初始化路由
	routers.Init(shell)
	shell.Run()
}
