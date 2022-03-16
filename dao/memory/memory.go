package memory

import (
	"github.com/ShangRui-hash/arp-spoofing-go/models"
)

var DataCh chan *models.HTTPPacket = make(chan *models.HTTPPacket, 2048)
