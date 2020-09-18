package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/abrie/mapatlapi/cmd/cli/command/server/handler"
)

func Run(args []string) {
	flagSet := flag.NewFlagSet("server", flag.ExitOnError)
	port := flagSet.Int("port", 0, "Listen for requests on this port.")
	flagSet.Parse(args)

	if *port == 0 {
		flagSet.Usage()
		os.Exit(2)
	}

	http.Handle("/geocoder", handler.Geocoder{})
	http.Handle("/location", handler.Location{})
	http.Handle("/places", handler.Places{})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Service started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
