package main

import "fmt"

func main() {

	tv1 := &SammysangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          false,
	}

	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    false,
	}

	useTV(&SammysangAdapter{sstv: tv1})
	fmt.Println("----------------------------")
	useTV(tv2)
}

func useTV(tv television) {
	tv.turnOn()
	tv.volumeUp()
	tv.volumeDown()
	tv.channelUp()
	tv.channelDown()
	tv.goToChannel(68)
	tv.turnOff()
}
