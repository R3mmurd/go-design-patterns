package main

import (
	"fmt"
	"math/rand"
)

// Miner is a concrete state machine
type Miner struct {
	StateMachine
	Energy int
}

// MinerWorking is a concrete state for a miner
type MinerWorking struct {
	context        *Miner
	requiredEnergy int
}

func (mw *MinerWorking) Enter() {
	fmt.Println("I am ready to work!!")
}

func (mw *MinerWorking) Handle() {
	if rand.Float64() < 0.2 {
		fmt.Println("Not lucky. I keep trying!!!")
	} else {
		fmt.Println("I got a good gold nugget")
	}

	mw.context.Energy -= mw.requiredEnergy

	if mw.context.Energy <= 0 {
		mw.context.Energy = 0
		mw.context.ChangeState("drinking")
	}
}

func (mw *MinerWorking) Exit() {
	fmt.Println("I am exhausted, I need a drink!")
}

// MinerDrinking is a concrete state for a miner
type MinerDrinking struct {
	context         *Miner
	recoveredEnergy int
}

func (md *MinerDrinking) Enter() {
	fmt.Println("Just I need this!!")
}

func (md *MinerDrinking) Handle() {
	fmt.Println("Ahhhh!! Good drink!!!")

	md.context.Energy += md.recoveredEnergy

	if md.context.Energy == 10 {
		md.context.ChangeState("working")
	}
}

func (md *MinerDrinking) Exit() {
	fmt.Println("Yeah! I quenched my thirst!")
}
