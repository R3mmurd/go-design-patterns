package main

// observable is the interface that need to be observed
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// DataSubject is the concrete observable
type DataSubject struct {
	observers []observer
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(obs observer) {
	ds.observers = append(ds.observers, obs)
}

func (ds *DataSubject) unregisterObserver(obs observer) {
	var newObservers []observer

	for _, o := range ds.observers {
		if obs.getID() != o.getID() {
			newObservers = append(newObservers, o)
		}
	}

	ds.observers = newObservers
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}
