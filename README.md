# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided by the City of Atlanta. Users may enter a street address and recieve information about the area, such as council district, nearby parks, transport, tax and emergency services. The service is useful, but lacks an API.

# MapAtlApi

This repository provides an (unofficial) API to [MapATL](https://egis.atlantaga.gov/app/home/index.html). It has three aspects:

- [Commandline interface](#cli)
- [Microservice server](#microservice)
- [Go module](#go-module)

### CLI

The CLI exposes the API as commandline arguments. For example:

`docker run abriedev/mapatlapi` to see the help screen.

1. `docker run abriedev/mapatlapi geocoder -address="55 Trinity Ave SW"` returns JSON with candidate Ref_IDs.
2. `docker run abriedev/mapatlapi location -id=490131` Uses Ref_ID to retrieve the location.
3. `docker run abriedev/mapatlapi places -id=490131 -category=PL_PARKS` Returns parks nearby.

### Microservice

The app may also be run as a webserver, for example:

`docker run -p 8888:8000 abriedev/mapatlapi:latest server -port 8000`

Exposes a server under http://localhost:8888 with these routes:

1. `/geocoder?address=55+Trinity+Ave+SW`
2. `/locations?id=490131`
3. `/places?id=490131&category=PL_PARKS`

### Go Module

To import the module:
`import "https://github.com/abrie/mapatlapi"`

The module exposes three functions mirroring the commandline/microservice interface
```go
func SearchByAddress(ctx context.Context, address string, maxLocations int64) (*geocoder.Response, error)
func FetchLocation(ctx context.Context, refId int) (*location.Response, error)
func FetchPlaces(ctx context.Context, refId int, category string) (*places.Response, error)
```

## Build Instructions

`make all` to test, build, and containerize.

## Technical Background

Internally, the [MapATL website](https://egis.atlantaga.gov/app/home/index.html) performs three interesting RPC calls.

1. _Geocoder_:
	- Input: An address, ex. "123 Peachtree St."
	- Output: A list of candidate locations, each with a `Ref_ID` number.
2. _Location_:
	- Input: A `Ref_ID` number
	- Output: Information about the location. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/point/response.go) for details.
3. _Places_ (_of interest_):
	- Input: A `Ref_ID` and a `Category`, which can be one of two values: `"PL_PARKS"` or `"TRANS_MARTA_RAIL_STATIONS"`.
	- Output: A list of places matching the category. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/places/response.go) for details.

Unfortunately some of these calls are POST's and therefore subject to CORS restrictions.


