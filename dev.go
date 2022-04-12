//go:generate protoc --go_out=. --go-grpc_out=. proto/bootstrap.proto
package main

import (
	"github.com/Meduzz/devserver/commands"
)

func main() {
	err := commands.Root.Execute()

	if err != nil {
		panic(err)
	}
}
