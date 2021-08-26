package main

import "fmt"

// newspaper is a specialization of publication
type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
