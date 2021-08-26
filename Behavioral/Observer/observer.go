package main

import "fmt"

// observer is the interface to listen an observable
type observer interface {
	getID() string
	onUpdate(data string)
}

// DataListener is the concrete observer
type DataListener struct {
	Name string
}

func (dl *DataListener) getID() string {
	return dl.Name
}

func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener", dl.Name, "got data change:", data)
}
