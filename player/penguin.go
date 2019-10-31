package player

import (
	"net"
)

type Penguin struct {
	Conn net.Conn
	Port string

	ID int
	Name string
	Coins int
	Room, X, Y int
	//Clothing struct	{
		//Color, Head, Face, Neck, Body, Hands, Feet, Pin, Background int
	//}

}

func getPenguins(){
	
}