package point

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
	flagSet := flag.NewFlagSet("point", flag.ExitOnError)
	refId := flagSet.Int("id", 0, "Id of the point to retrieve")

	flagSet.Parse(args)

	if *refId == 0 {
		flagSet.Usage()
		fmt.Printf("\nError: Id must be a number greater than zero.")
		os.Exit(2)
	}

	result, err := mapatlapi.FetchPoint(context.Background(), *refId)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)

}

func Help() {
}
