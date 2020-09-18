# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided by the City of Atlanta. Through an HTML interface, a user can enter a street address and recieve information about the location such as council district, nearby parks, transport, tax and emergency services. This repository is an attempt to provide an API to the service.

## Technical Background

Internally, the [MapATL website](https://egis.atlantaga.gov/app/home/index.html) takes an address and performs three interesting RPC calls.

1. _Geocoder_:
	- Input: An address, ex. "123 Peachtree St."
	- Output: A list of candidate locations, each with a `Ref_ID` number.
2. _Location_:
	- Input: A `Ref_ID` number
	- Output: Information about the location. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/point/response.go) for details.
3. _Places_ (_of interest_):
	- Input: A `Ref_ID` and a `Category`, which can be one of two values: `"PL_PARKS"` or `"TRANS_MARTA_RAIL_STATIONS"`.
	- Output: A list of places matching the category. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/places/response.go) for details.

Unfortunately some of these calls are POST's and subject to CORS restrictions; which limits the usefulness for web apps.

## Unofficial CLI

Included in this repository is a CLI application packaged as a Docker container. Run it like this:

`docker run docker.pkg.github.com/abrie/mapatlapi/mapatlcli:latest`

The CLI may also be run as a webserver, easily deployable as a microservice:

`docker run docker.pkg.github.com/abrie/mapatlapi/mapatlservice:latest server -port XXXX`

The following routes mirror the RPC calls described above:

- `/geocoder?address=url-encoded-address-parameter-here`
- `/locations?id=ref_id_here`
- `/places?id=ref_id_here&category=category_here`

## Unofficial API

This repository is an importable Go module.

`import "https://github.com/abrie/mapatlapi"`

It provides three methods, essentially lightweight wrappers around the RPC calls:

```go
func SearchByAddress(ctx context.Context, address string) (*geocoder.Response, error)
func FetchLocation(ctx context.Context, refId int) (*location.Response, error)
func FetchPlaces(ctx context.Context, refId int, category string) (*places.Response, error)
```

