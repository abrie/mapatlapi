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

func Run(args []string) {
	flagSet := flag.NewFlagSet("geocoder", flag.ExitOnError)
	address := flagSet.String("address", "", "Address to geocode, for example '123 Peachtree St.'")

	flagSet.Parse(args)

	if *address == "" {
		fmt.Printf("Error: specify an address to geocode.\n")
		Help()
		os.Exit(2)
	}

	result, err := mapatlapi.SearchByAddress(context.Background(), *address)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(*result)

}

func Help() {
}
