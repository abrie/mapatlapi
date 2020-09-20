package places

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
	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'id' parameter.")
		return
	}

	categories, ok := r.URL.Query()["category"]
	if !ok || len(categories) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'category' parameter.")
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "'%s' is not parsable into an int.", ids[0])
		return
	}

	result, err := mapatlapi.FetchPlaces(context.Background(), id, categories[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to fetch places: %s", err.Error())
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal places response into JSON: %s", err.Error())
		fmt.Fprintf(w, "Internal API called failed.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if size, err := w.Write(response); err != nil {
		log.Printf("Failed write location response: Wrote %d bytes, received error %s", size, err.Error())
		return
	}

}
