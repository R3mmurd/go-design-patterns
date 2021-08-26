package main

// iPublication is an interface for any type of publication.
type iPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	getName() string
	getPages() int
	getPublisher() string
}

// publication is a base type for publications.
type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p *publication) setPublisher(publisher string) {
	p.publisher = publisher
}

func (p *publication) getName() string {
	return p.name
}

func (p *publication) getPages() int {
	return p.pages
}

func (p *publication) getPublisher() string {
	return p.publisher
}
