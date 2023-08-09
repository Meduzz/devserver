package main

import "github.com/Meduzz/devserver/commands"

func main() {
	root := commands.Root

	err := root.Execute()

	if err != nil {
		panic(err)
	}
}
