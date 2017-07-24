package main

import "fmt"
import "./players"

var cpuPlayerList []players.ComputerPlayer
var networkPlayerList []players.NetworkPlayer
var playerList []players.Player

func main() {
	quitChan := make(chan struct{})
	go serverFunc(quitChan)
	go gameLoop()
	fmt.Println("aaaaaaaaa")
}