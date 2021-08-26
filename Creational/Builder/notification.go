package main

// Notification is the data type that will be created by the builder.
type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}
