package main

func main() {
	listener1 := &DataListener{
		Name: "Listener 1",
	}

	listener2 := &DataListener{
		Name: "Listener 2",
	}

	subject := &DataSubject{}

	subject.registerObserver(listener1)
	subject.registerObserver(listener2)

	subject.ChangeItem("First change")

	subject.unregisterObserver(listener1)

	subject.ChangeItem("Second change")
}
