package memory

import "github.com/ShangRui-hash/arp-spoofing-go/models"

var (
	loots = make([]*models.Loot, 0)
)

func AddLoot(loot *models.Loot) {
	loots = append(loots, loot)
}

func GetAllLoot() []*models.Loot {
	return loots
}

func ClearAllLoot() {
	loots = make([]*models.Loot, 0)
	return
}
