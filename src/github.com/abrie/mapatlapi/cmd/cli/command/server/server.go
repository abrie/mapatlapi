package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/cli/command/geocoder"
	"github.com/abrie/mapatlapi/cmd/cli/command/location"
	"github.com/abrie/mapatlapi/cmd/cli/command/places"
)

func Run(args []string) {
	flagSet := flag.NewFlagSet("server", flag.ExitOnError)
	port := flagSet.Int("port", 0, "Listen for requests on this port.")
	flagSet.Parse(args)

	if *port == 0 {
		flagSet.Usage()
		os.Exit(2)
	}

	http.HandleFunc("/geocoder", geocoder.ServeHTTP)
	http.HandleFunc("/location", location.ServeHTTP)
	http.HandleFunc("/places", places.ServeHTTP)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Service started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Help() {
}
