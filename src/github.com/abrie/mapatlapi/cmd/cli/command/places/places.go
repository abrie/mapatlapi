package places

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

/* TODO: Move Category definitions to the api */
var categories map[string]bool

func init() {
	categories = map[string]bool{
		"PL_PARKS":                  true,
		"TRANS_MARTA_RAIL_STATIONS": true}
}

func Run(api mapatlapi.Config, args []string) {
	flagSet := flag.NewFlagSet("places", flag.ExitOnError)
	refId := flagSet.Int("id", 0, "Id of the location to retrieve")
	category := flagSet.String("category", "", "Category of places of interest corresponding to locator")

	flagSet.Parse(args)

	if *refId == 0 {
		flagSet.Usage()
		fmt.Printf("\nError: Id must be a number greater than zero.")
		os.Exit(2)
	}

	if _, found := categories[*category]; !found {
		flagSet.Usage()
		fmt.Printf("\nError: Valid categories are:\n")
		for key := range categories {
			fmt.Printf("\t%s\n", key)
		}
		os.Exit(2)
	}

	result, err := api.FetchPlaces(context.Background(), *refId, *category)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)

}

func Help() {
}
