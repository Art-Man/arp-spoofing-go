package server

import (
	"encoding/json"
	"log"

	"github.com/ShangRui-hash/arp-spoofing-go/dao/memory"

	"github.com/gorilla/websocket"
)

//writeTo 向Conn中写
func writeTo(conn *websocket.Conn) {
	for data := range memory.DataCh {
		datastr, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, datastr); err != nil {
			log.Println("conn write failed,err:", err)
			return
		}
	}
}

//readFrom 从Conn中读
func readFrom(conn *websocket.Conn) {
	return
}
