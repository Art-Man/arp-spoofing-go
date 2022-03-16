package logic

import (
	"fmt"

	"github.com/ShangRui-hash/arp-spoofing-go/dao/memory"
)

//ShowLoot 展示所有战利品
func ShowLoot() error {
	loots := memory.GetAllLoot()
	if len(loots) == 0 {
		return nil
	}
	for _, loot := range loots {
		fmt.Printf("%s:%s->%s:%s [%s]\n", loot.SrcIP, loot.SrcPort, loot.DstIP, loot.DstPort, loot.Keyword)
		fmt.Println(loot.Payload)
	}
	return nil
}

//ClearLoot 清除所有战利品
func ClearLoot() error {
	memory.ClearAllLoot()
	return nil
}
