package geocoder

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

import (
	"github.com/abrie/mapatlapi"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addresses, ok := r.URL.Query()["address"]

	if !ok || len(addresses) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Exactly 1 'address' parameter required.")
		return
	}

	addressArg := addresses[0]

	var maxLocationsArg int64 = 6

	maxLocations, ok := r.URL.Query()["maxLocations"]
	if ok {
		if len(maxLocations) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Exactly 0 or 1 'maxLocations' parameter required.")
			return
		}

		val, err := strconv.ParseInt(maxLocations[0], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "maxLocations must be a number. ")
			return
		}
		maxLocationsArg = val
	}

	result, err := mapatlapi.SearchByAddress(context.Background(), addressArg, maxLocationsArg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to search address: %s", err)
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal address response into JSON: %s", err)
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		log.Printf("Failed write address response: %s", err)
		return
	}

}
