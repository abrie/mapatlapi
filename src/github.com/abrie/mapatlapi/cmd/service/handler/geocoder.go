package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

import (
	"github.com/abrie/mapatlapi"
)

type Geocoder struct {
}

func (geocoder Geocoder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addresses, ok := r.URL.Query()["address"]

	if !ok || len(addresses) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'address' parameter.")
		return
	}

	result, err := mapatlapi.SearchByAddress(context.Background(), addresses[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to search address: %s", err.Error())
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal address response into JSON: %s", err.Error())
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if size, err := w.Write(response); err != nil {
		log.Printf("Failed write address response: Wrote %d bytes, received error %s", size, err.Error())
		return
	}

}
