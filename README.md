# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided to the public by the Atlanta government. It provides an HTML interface for searching by street address; returning information about the neighborhood such as districts, parks, transport, tax and services. It's very useful, but unfortunately the service provides no offical API.

## Unofficial API

This repository provides a Go module for programmatic interaction with MapATL.


### Example Code

```go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

import (
	"github.com/abrie/mapatlapi"
)

func main() {

	const oneLineAddress = "123 Peachtree St."

	/* Invoke SearchByAddress with a string containing the address to search */
	result, err := mapatlapi.SearchByAddress(oneLineAddress)
	if err != nil {
		log.Fatal(err)
	}

	/* If the one line address is valid, then API will return a list of candidates.
	   If the candidate list is empty then the address either doesn't exist,
		 or needs to be more specific.
	*/
	if len(result.Candidates) == 0 {
		fmt.Printf("No candidates found for address \"%s\"\n", oneLineAddress)
		os.Exit(0)
	}

	/* Ideally the returned result will contain only one candidate address,
	   but "123 Peachtree St" is vague and contains several possibilities.
	*/
	for _, candidate := range result.Candidates {
		fmt.Printf("Candidate: %s\n", candidate.Address)

		/* Each candidate contains a Ref_ID field which must be used to retrieve related data.*/
		result, err := mapatlapi.FetchRecord(candidate.Attributes.Ref_ID)
		if err != nil {
			log.Fatal(err)
		}

		/* If the Ref_ID is invalid, then no results will be returned. */
		if len(result.Records) == 0 {
			fmt.Printf("No records founds where Ref_ID=%d.\n", candidate.Attributes.Ref_ID)
			os.Exit(0)
		}

		/* The MapATL service returns an array of results, but most likely the length will be 1. */
		for _, record := range result.Records {
			/* We just dump out the entire record as JSON. */
			bytes, err := json.MarshalIndent(record, "", " ")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\n", string(bytes))
		}
	}
}
```
