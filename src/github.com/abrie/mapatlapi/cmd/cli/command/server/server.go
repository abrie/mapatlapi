package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/abrie/mapatlapi"
	"github.com/abrie/mapatlapi/cmd/cli/command/geocoder"
	"github.com/abrie/mapatlapi/cmd/cli/command/location"
	"github.com/abrie/mapatlapi/cmd/cli/command/places"
)

func Run(api mapatlapi.Config, args []string) {
	flagSet := flag.NewFlagSet("server", flag.ExitOnError)
	port := flagSet.Int("port", 0, "Listen for requests on this port.")
	flagSet.Parse(args)

	if *port == 0 {
		flagSet.Usage()
		os.Exit(2)
	}

	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/geocoder", geocoder.HandlerFunc(api))
	http.HandleFunc("/location", location.HandlerFunc(api))
	http.HandleFunc("/places", places.HandlerFunc(api))

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Service started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Help() {
}
