package main

import "fmt"

func main() {
	// Iterating by using a callback
	lib.IterateBooks(bookCallback)

	// Iterating by using an anonymous callback
	lib.IterateBooks(func(book Book) error {
		fmt.Println(book.Title, "by", book.Author)
		return nil
	})

	// Iterating by using an iterator interface
	it := lib.createIterator()
	for it.hasNext() {
		book := it.getNext()
		fmt.Println(book)
	}
}

func bookCallback(book Book) error {
	fmt.Println("Book title:", book.Title)
	return nil
}
