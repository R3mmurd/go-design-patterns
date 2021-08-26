package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()

	builder.SetTitle("New Notification")
	builder.SetSubTitle("This is a subtitle")
	builder.SetIcon("icon.png")
	builder.SetImage("image.png")
	builder.SetPriority(3)
	builder.SetMessage("This is a basic Notification with some text in it")
	builder.SetType("alert")

	notification, err := builder.Build()

	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v\n", notification)
	}
}
