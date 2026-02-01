package main

import (
	"log"

	"github.com/Meduzz/commando"
	_ "github.com/Meduzz/devserver/cmd"
)

func main() {
	err := commando.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
