# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided by the Atlanta government. It provides an HTML interface for searching by street address; returning information such as council districts, parks, transport, tax and emergency services. Unfortunately the service provides no offical API.

## Technical Background

Internally, the [MapATL website](https://egis.atlantaga.gov/app/home/index.html) takes an address and performs three interesting RPC calls. This API is a lightweight wrapper around those calls.

1. _Geocoder_:
	- Input: An address, ex. "123 Peachtree St."
	- Output: A list of candidate locations, each with a `Ref_ID` number.
2. _Location_:
	- Input: A `Ref_ID` number
	- Output: Information about the location. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/point/response.go) for details.
3. _Places_ (_of interest_):
	- Input: A `Ref_ID` and a `Category`, which can be one of two values: `"PL_PARKS"` or `"TRANS_MARTA_RAIL_STATIONS"`.
	- Output: A list of places matching the category. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/places/response.go) for details.

## Unofficial API

This repository provides a Go API for programmatic access to the site. It provides three methods:

```go
func SearchByAddress(ctx context.Context, address string) (*geocoder.Response, error)
func FetchLocation(ctx context.Context, refId int) (*location.Response, error)
func FetchPlaces(ctx context.Context, refId int, category string) (*places.Response, error)
```

## Unofficial CLI

Included in this repository is a CLI application using the Go API. It's been wrapped as a Docker container. Run it like this:

`docker run docker.pkg.github.com/abrie/mapatlapi/mapatlcli:latest`

Run the command without parameters for help.

## Unofficial Microservice

There also a minimal server available which exposes the API functions through GETable handlers. Run it like this:

`docker run docker.pkg.github.com/abrie/mapatlapi/mapatlservice:latest -port XXXX`

The microservice listens on the specified port, and responsds to GET requests as follows:

- `/geocoder?address=url-encoded-address-parameter-here`
- `/locations?id=ref_id_here`
- `/places?is=ref_id_here&category=category_here`

