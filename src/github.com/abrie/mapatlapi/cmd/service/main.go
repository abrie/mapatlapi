package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/service/handler"
)

func main() {
	port := flag.Int("port", 0, "Listen for requests on this port.")
	flag.Parse()

	if *port == 0 {
		flag.Usage()
		os.Exit(2)
	}

	http.Handle("/geocoder", handler.Geocoder{})
	http.Handle("/location", handler.Location{})
	http.Handle("/places", handler.Places{})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Service started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
