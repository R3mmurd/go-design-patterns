package main

import "fmt"

// BookType represents the type of book
type BookType int

// Predefined Book types
const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

// Book represents data about a book
type Book struct {
	Title     string
	Author    string
	PageCount int
	Type      BookType
}

// Library represents the collection of books
type Library struct {
	Collection []Book
}

// IterateBooks calls the given callback for each book in the collection
func (l *Library) IterateBooks(f func(Book) error) {
	var err error

	for _, book := range l.Collection {
		err = f(book)

		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

// Example of a book collection
var lib *Library = &Library{
	Collection: []Book{
		{
			Title:     "War and Peace",
			Author:    "Leo Tolstoy",
			PageCount: 864,
			Type:      HardCover,
		},
		{
			Title:     "Crime and Punishment",
			Author:    "Leo Tolstoy",
			PageCount: 1225,
			Type:      SoftCover,
		},
		{
			Title:     "Brave New World",
			Author:    "Aldous Huxley",
			PageCount: 325,
			Type:      PaperBack,
		},
		{
			Title:     "Catcher in the Rye",
			Author:    "J.D. Salinger",
			PageCount: 206,
			Type:      HardCover,
		},
		{
			Title:     "To Kill a Mockingbird",
			Author:    "Harper Lee",
			PageCount: 399,
			Type:      PaperBack,
		},
		{
			Title:     "The Grapes of Wrath",
			Author:    "John Steinbeck",
			PageCount: 464,
			Type:      HardCover,
		},
		{
			Title:     "Wuthering Heights",
			Author:    "Emily Bronte",
			PageCount: 288,
			Type:      EBook,
		},
	},
}
