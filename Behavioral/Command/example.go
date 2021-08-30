package main

func main() {
	npc := &NPC{
		name: "NPC 1",
	}
	npc.commands = map[string]iCommand{
		"hello": &CommandGreet{
			context: npc,
		},
		"info": &CommandGiveInfo{},
		"bye":  &CommandDimiss{},
	}

	npc.ListenAndAct("hello")
	npc.ListenAndAct("info")
	npc.ListenAndAct("bye")
}
