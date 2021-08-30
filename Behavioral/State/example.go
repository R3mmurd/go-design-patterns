package main

import (
	"log"
	"time"
)

func main() {
	miner := &Miner{
		Energy: 10,
	}
	miner.states = map[string]iState{
		"working": &MinerWorking{
			context:        miner,
			requiredEnergy: 2,
		},
		"drinking": &MinerDrinking{
			context:         miner,
			recoveredEnergy: 1,
		},
	}
	miner.currentState = "working"

	for i := 0; i < 100; i++ {
		err := miner.Exec()

		if err != nil {
			log.Fatalln(err)
			break
		}

		time.Sleep(200 * time.Millisecond)
	}
}
