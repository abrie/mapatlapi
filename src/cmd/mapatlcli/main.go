package main

import "github.com/abrie/mapatlapi"
import "log"
import "fmt"
import "flag"
import "os"

func main() {
	addressPtr := flag.String("address", "", "Single line address")
	flag.Parse()

	if *addressPtr == "" {
		fmt.Printf("mapatlcli: Specify an address to search, for example:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: mapatlcli -address=\"123 Peachtree St.\"\n")
		os.Exit(1)
	}

	result, err := mapatlapi.SearchByAddress(*addressPtr)
	if err != nil {
		log.Fatal(err)
	}

	for _, candidate := range result.Candidates {
		fmt.Printf("%s\t%d\t%f\n", candidate.Address, candidate.Attributes.Ref_ID, candidate.Attributes.Score)
	}

	for _, candidate := range result.Candidates {
		result, err := mapatlapi.FetchRecord(candidate.Attributes.Ref_ID)
		if err != nil {
			log.Fatal(err)
		}

		for _, record := range *result {
			fmt.Printf("%s\t%s\n", record.COUNCIL_DIST, record.LABEL)
		}
	}
}
