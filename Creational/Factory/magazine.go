package main

import "fmt"

// magazine is a specialization of publication
type magazine struct {
	publication
}

func (m magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", m.name)
}

func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
