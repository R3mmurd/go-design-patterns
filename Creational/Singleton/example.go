package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getLoggerInstance()
			wg.Done()
		}()
	}

	wg.Wait()

	log := getLoggerInstance()

	log.SetLogLevel(1)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(2)
	log.Log("This is a log message")

	log = getLoggerInstance()
	log.SetLogLevel(3)
	log.Log("This is a log message")
}
