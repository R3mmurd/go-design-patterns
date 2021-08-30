package main

import (
	"fmt"
)

// NPC is an example of a game NPC
type NPC struct {
	name     string
	commands map[string]iCommand
}

func (npc *NPC) ListenAndAct(text string) {
	command, ok := npc.commands[text]

	if !ok {
		fmt.Println("I don't understand what do you need!")
		return
	}

	command.Execute()
}

type CommandGreet struct {
	context *NPC
}

func (cg *CommandGreet) Execute() {
	fmt.Printf("Hello! My name is %s. How could I help you?\n", cg.context.name)
}

type CommandGiveInfo struct{}

func (cgi *CommandGiveInfo) Execute() {
	fmt.Println("Follow your heart and you'll find your destiny!")
}

type CommandDimiss struct{}

func (cd *CommandDimiss) Execute() {
	fmt.Println("Good bye! Nice to meet you!")
}
