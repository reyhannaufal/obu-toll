package main

import (
	"fmt"
	"log"
	"math/rand"
	"obu-toll/types"
	"time"

	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

func main() {
	obuIDS := geenerateObuIDS(20)
	fmt.Println(obuIDS)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := genCoords()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(5 * time.Second)
	}

}

func geenerateObuIDS(n int) []int {
	ids := make([]int, n)
	for j := 0; j < n; j++ {
		ids[j] = rand.Intn(999999)
	}
	return ids
}

func genCoords() (float64, float64) {
	return genLatLong(), genLatLong()
}

func genLatLong() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}
