package main

import (
	"log"
	"net/http"
	"obu-toll/types"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

func main() {
	rcv := NewDataReceiver()
	http.HandleFunc("/ws", rcv.handleWS)
	http.ListenAndServe(":30000", nil)

}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiver()

}

func (dr *DataReceiver) wsReceiver() {
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Fatal(err)
			continue
		}
		dr.msgch <- data
	}
}
