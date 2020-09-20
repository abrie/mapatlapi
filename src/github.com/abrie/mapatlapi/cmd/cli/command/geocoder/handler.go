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

	if !ok || len(addresses) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'address' parameter.")
		return
	}

	addressArg := addresses[0]

	maxLocations, ok := r.URL.Query()["maxLocations"]
	if !ok || len(maxLocations) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "maxLocations must not be specified more than once.")
		return
	}

	var maxLocationsArg int64
	if len(maxLocations) == 1 {
		val, err := strconv.ParseInt(maxLocations[0], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "maxLocations must be a number. ")
			return
		}
		maxLocationsArg = val
	} else {
		maxLocationsArg = 6
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
