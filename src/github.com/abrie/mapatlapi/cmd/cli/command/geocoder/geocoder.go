package geocoder

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

import (
	"github.com/abrie/mapatlapi"
	"github.com/abrie/mapatlapi/cmd/cli/utils"
)

func Run(config mapatlapi.Config, args []string) {
	flagSet := flag.NewFlagSet("geocoder", flag.ExitOnError)
	address := flagSet.String("address", "", "Address to geocode, for example '123 Peachtree St.'")
	maxLocations := flagSet.Int64("maxLocations", 6, "Maximum returned candidate locations.")

	flagSet.Parse(args)

	if *address == "" {
		fmt.Printf("Error: specify an address to geocode.\n")
		Help()
		os.Exit(2)
	}

	result, err := config.SearchByAddress(context.Background(), *address, *maxLocations)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(*result)

}

func Help() {
}
