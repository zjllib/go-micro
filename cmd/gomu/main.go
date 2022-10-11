package main

import (
	"github.com/zjllib/go-micro/cmd/gomu/cmd"

	// register commands
	_ "github.com/zjllib/go-micro/cmd/gomu/cmd/cli"
)

func main() {
	cmd.Run()
}
