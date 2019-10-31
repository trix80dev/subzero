package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "handlers"
    "player"
    "util"
)

var connection int

func main() {

    fmt.Println(" __     __          _                    \r\n" + 
    " \\ \\   / /         | |                   \r\n" + 
    "  \\ \\_/ /   _   _  | | __   ___    ____  \r\n" + 
    "   \\   /   | | | | | |/ /  / _ \\  |  _ \\ \r\n" + 
    "    | |    | |_| | |   <  | (_) | | | | |\r\n" + 
    "    |_|     \\__,_| |_|\\_\\  \\___/  |_| |_|\r\n")

    port1 := "6112"
    listener, err := net.Listen("tcp", "localhost:"+port1)

    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    port2 := "6113"
    listener2, err := net.Listen("tcp", "localhost:"+port2)

    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    util.StartDatabase()

    go acceptLoop(listener, port1)
    acceptLoop(listener2, port2)
}

func acceptLoop(listener net.Listener, port string){
    defer listener.Close()
    fmt.Println("Listening on port", port)
    for {

        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        fmt.Println("Client connected to port", port)
        penguin := player.Penguin{Conn:conn, Port: port, ID: 101, Name: "chuh", Coins: 100, Room: 100, X: 100, Y:100}
        go handleRequest(penguin)
    }
}

func handleRequest(penguin player.Penguin) {

    connection++

    number := connection

  for{

    data := make([]byte, 1024)

    read, err := penguin.Conn.Read(data)
  
    //fmt.Println(number)

    if read > 0 {

    if err != nil {
        fmt.Println("Error reading:", err.Error())
        }

        packet := string(data)
        packet = strings.TrimRight(packet, "\x00")

        fmt.Println("IN:", packet, number)

        if strings.HasPrefix(packet, "<") {
            handlers.HandleXMLPacket(penguin, packet)
        } else if strings.HasPrefix(packet, "%"){
            handlers.HandleXTPacket(penguin, packet)
        } else {
            fmt.Println("Unknown packet type")
        }

        }
    }
}