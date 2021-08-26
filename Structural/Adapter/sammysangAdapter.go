package main

type SammysangAdapter struct {
	sstv *SammysangTV
}

func (sa *SammysangAdapter) volumeUp() int {
	vol := sa.sstv.getVolume() + 1
	sa.sstv.setVolume(vol)
	return vol
}

func (sa *SammysangAdapter) volumeDown() int {
	vol := sa.sstv.getVolume() - 1
	sa.sstv.setVolume(vol)
	return vol
}

func (sa *SammysangAdapter) channelUp() int {
	ch := sa.sstv.getChannel() + 1
	sa.sstv.setChannel(ch)
	return ch
}

func (sa *SammysangAdapter) channelDown() int {
	ch := sa.sstv.getChannel() - 1
	sa.sstv.setChannel(ch)
	return ch
}

func (sa *SammysangAdapter) turnOn() {
	sa.sstv.setOnState(true)
}

func (sa *SammysangAdapter) turnOff() {
	sa.sstv.setOnState(false)
}

func (sa *SammysangAdapter) goToChannel(ch int) {
	sa.sstv.setChannel(ch)
}
