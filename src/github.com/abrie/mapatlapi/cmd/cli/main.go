package main

import (
	"fmt"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/cli/command/geocoder"
	"github.com/abrie/mapatlapi/cmd/cli/command/places"
	"github.com/abrie/mapatlapi/cmd/cli/command/point"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("mapatlcli is a cli interface to MapATL.\n")
		fmt.Printf("\nusage: mapatlcli <command> [params]\n")
		fmt.Printf("\nCommands are:\n")
		fmt.Printf("\tgeocoder\tTransforms an address to a point ID.\n")
		fmt.Printf("\tpoint\t\tLoads the data associated with a point ID.\n")
		fmt.Printf("\tplaces\t\tReturns places of interest near to a point ID.\n")
		os.Exit(2)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "geocoder":
		geocoder.Run(args)
	case "point":
		point.Run(args)
	case "places":
		places.Run(args)
	default:
		fmt.Printf("Unkown command '%s'.\n", command)
		os.Exit(2)
	}

	os.Exit(0)
}
