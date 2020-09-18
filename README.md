# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided by the City of Atlanta. Users may enter a street address and recieve information about the area, such as council district, nearby parks, transport, tax and emergency services. The service is useful, but lacks an API.

# About MapAtlApi

This repository provides a microservice/cli and Go module for programmatic interaction with [MapATL](https://egis.atlantaga.gov/app/home/index.html).

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

### Unofficial CLI

The CLI mirrors the above RPC methods as commandline arguments. For example:

- `docker run abriedev/mapatlapi` to see the help screen.
1. `docker run abriedev/mapatlapi geocoder -address="55 Trinity Ave SW"` returns candidate Ref_IDs.
2. `docker run abriedev/mapatlapi location -id=490131"` Uses Ref_ID to retrieve the location.
3. `docker run abriedev/mapatlapi places -id=490131 -category=PL_PARKS` Uses the Ref_ID to return nearby parks.

The CLI may also be run as a webserver:

`docker run abriedev/mapatlapi server -port XXXX`

The following routes mirror the CLI commands described above:

1. `/geocoder?address=url-encoded-address-parameter-here`
2. `/locations?id=ref_id_here`
3. `/places?id=ref_id_here&category=category_here`

## Unofficial API

This repository is an importable Go module.

`import "https://github.com/abrie/mapatlapi"`

It provides three methods, essentially lightweight wrappers around the RPC calls:

```go
func SearchByAddress(ctx context.Context, address string) (*geocoder.Response, error)
func FetchLocation(ctx context.Context, refId int) (*location.Response, error)
func FetchPlaces(ctx context.Context, refId int, category string) (*places.Response, error)
```

