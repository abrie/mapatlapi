package main

import "github.com/abrie/mapatlapi"
import "log"
import "fmt"
import "flag"
import "os"
import "encoding/json"

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

	if len(result.Candidates) == 0 {
		fmt.Printf("No candidates found for address \"%s\"\n", *addressPtr)
		os.Exit(0)
	}

	for _, candidate := range result.Candidates {
		fmt.Printf("%s\t%d\t%f\n", candidate.Address, candidate.Attributes.Ref_ID, candidate.Attributes.Score)
	}

	for _, candidate := range result.Candidates {
		result, err := mapatlapi.FetchRecord(candidate.Attributes.Ref_ID)
		if err != nil {
			log.Fatal(err)
		}

		if len(result.Records) == 0 {
			fmt.Printf("No records founds where Ref_ID=%d.\n", candidate.Attributes.Ref_ID)
			os.Exit(0)
		}

		for _, record := range result.Records {
			bytes, err := json.MarshalIndent(record, "", " ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", string(bytes))
		}
	}
}
