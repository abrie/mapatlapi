package main

import "github.com/abrie/mapatlapi"
import "log"
import "fmt"
import "flag"

func main() {
	addressPtr := flag.String("address", "", "Single line address")
	flag.Parse()

	if *addressPtr == "" {
		fmt.Printf("mapatlcli: Need an address to search. Specify as follows:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: mapatlcli -address=\"123 Peachtree St.\"\n")
		return
	}

	result, err := mapatlapi.SearchByAddress(*addressPtr)
	if err != nil {
		log.Fatal(err)
	}

	for _, candidate := range result.Candidates {
		fmt.Printf("%s\t%d\t%f\n", candidate.Address, candidate.Attributes.Ref_ID, candidate.Attributes.Score)
	}
}
