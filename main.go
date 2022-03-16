package main

import (
	"log"

	"github.com/ShangRui-hash/arp-spoofing-go/logo"
	"github.com/ShangRui-hash/arp-spoofing-go/routers"
	"github.com/ShangRui-hash/arp-spoofing-go/settings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "阿弥陀佛 > ",
	})
	logo.Show(logo.LogoFile)
	//1.初始化配置
	if err := settings.Init(); err != nil {
		log.Println("settings.Init failed,err:", err)
		return
	}
	//2.连接redis
	// if err := redis.Init(); err != nil {
	// 	log.Println("redis init failed,err:", err)
	// 	return
	// }
	// defer redis.Close()
	// debug.Println("redis 数据库连接成功")
	routers.Init(shell)
	shell.Run()
}
