package main

import "fmt"

// newPublication is a factory function that creaters the specified publication
func newPublication(pubType string, name string, pages int, publisher string) (iPublication, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pages, publisher), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pages, publisher), nil
	}

	return nil, fmt.Errorf("no such publication type")
}
