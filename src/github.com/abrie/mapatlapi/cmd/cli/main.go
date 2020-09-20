package main

import (
	"fmt"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/cli/command/geocoder"
	"github.com/abrie/mapatlapi/cmd/cli/command/location"
	"github.com/abrie/mapatlapi/cmd/cli/command/places"
	"github.com/abrie/mapatlapi/cmd/cli/command/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("mapatlapi is an API interface to MapATL.\n")
		fmt.Printf("\nusage: mapatlcli <command> [params]\n")
		fmt.Printf("\nCommands are:\n")
		fmt.Printf("\tgeocoder\tTransforms an address to a point ID.\n")
		fmt.Printf("\tlocation\tLoads the data for the location associated with a point ID.\n")
		fmt.Printf("\tplaces\t\tReturns places of interest proximal to the location identified by the point ID.\n")
		fmt.Printf("\tserver\t\tStarts this app as a web service.\n")
		os.Exit(2)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "geocoder":
		geocoder.Run(args)
	case "location":
		location.Run(args)
	case "places":
		places.Run(args)
	case "server":
		server.Run(args)
	default:
		fmt.Printf("Unkown command '%s'.\n", command)
		os.Exit(2)
	}

	os.Exit(0)
}
