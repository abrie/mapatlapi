package main

import (
	"fmt"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/cli/command/geocoder"
	"github.com/abrie/mapatlapi/cmd/cli/command/records"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Requires a command.\n")
		os.Exit(2)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "geocoder":
		geocoder.Run(args)
	case "records":
		records.Run(args)
	default:
		fmt.Printf("Unkown command '%s'.\n", command)
		os.Exit(2)
	}

	os.Exit(0)
}
