package handlers

import (
	"fmt"
	"player"
)

func HandleOutput(penguin player.Penguin, packet string) {

	packet += "\u0000"
	penguin.Conn.Write([]byte(packet))
	
	fmt.Println("OUT:", packet)

}