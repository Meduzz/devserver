package main

import (
	"log"

	"github.com/Meduzz/commando"
)

func main() {
	err := commando.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
