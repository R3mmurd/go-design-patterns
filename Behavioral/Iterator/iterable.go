package main

// IteratorCollection describes an interface that creates a new iterator
type IteratorCollection interface {
	createIterator() iterator
}

// iterator is the interface for any iterable type
type iterator interface {
	hasNext() bool
	getNext() interface{}
}

// BookIterator is the concrete iterator
type BookIterator struct {
	current int
	books   []Book
}

func (it *BookIterator) hasNext() bool {
	return it.current < len(it.books)
}

func (it *BookIterator) getNext() interface{} {
	if !it.hasNext() {
		return nil
	}

	book := &it.books[it.current]
	it.current++
	return book
}
