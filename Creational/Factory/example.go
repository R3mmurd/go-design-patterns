package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Tyme", 50, "The Tymes")
	mag2, _ := newPublication("magazine", "Lyfe", 50, "Lyfe Inc")
	news1, _ := newPublication("newspaper", "The Herald", 50, "Heralders")
	news2, _ := newPublication("newspaper", "The Standard", 50, "Standarders")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Println("-----------------------------------")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Println("-----------------------------------")
}
