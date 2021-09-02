package main

import (
	"log"
	"time"
)

func main() {
	miner := newMiner(10, 2, 1)

	for i := 0; i < 100; i++ {
		err := miner.Exec()

		if err != nil {
			log.Fatalln(err)
			break
		}

		time.Sleep(200 * time.Millisecond)
	}
}
