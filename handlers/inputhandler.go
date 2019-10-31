package handlers

import (
	"fmt"
	"strings"
	"player"
	"util"
	"strconv"
)

type fn func (player.Penguin, string)

var XTPackets = map[string] fn {
	"j#js": HandleJoinServer,
}

func HandleXTPacket(penguin player.Penguin, packet string){

split := strings.Split(packet, "%")

fmt.Println(split[3])

for key, _ := range XTPackets {
	if(key == split[3]) {
		XTPackets[split[3]](penguin, packet)
		break;
	}
}

}

func HandleJoinServer(penguin player.Penguin, packet string){

	response := "%xt%js%-1%" + strconv.Itoa(penguin.ID) + "%0%0%0%"

	HandleOutput(penguin, response)

}

func HandleXMLPacket(penguin player.Penguin, packet string)	{

	var response string

	if strings.Contains(packet, "policy-file-request") {
		response = "<cross-domain-policy><allow-access-from domain='*' to-ports='" + penguin.Port + "'/></cross-domain-policy>"
	}	else if strings.Contains(packet, "body action='verChk'") {
		response = "<msg t='sys'><body action='apiOK' r='0'></body></msg>"
	}	else if strings.Contains(packet, "body action='rndK'") {
		response = "<msg t='sys'><body action='rndK' r='-1'><k>" + util.GenerateRandomKey() + "</k></body></msg>"
	}	else if strings.Contains(packet, "body action='login'") {
		response = "%xt%l%" + strconv.Itoa(penguin.ID) + "%" + util.GenerateRandomKey() + "%%50%" + penguin.Name + "%"
	}	else {
		fmt.Println("packet not handled")
		return
	}

	HandleOutput(penguin, response)

}