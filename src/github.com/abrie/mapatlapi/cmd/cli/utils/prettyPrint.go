package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(obj interface{}) {
	bytes, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(bytes))
}
