# About MapATL

[MapATL](https://egis.atlantaga.gov/app/home/index.html) is a web service provided by the Atlanta government. It provides an HTML interface for searching by street address; returning information such as council districts, parks, transport, tax and emergency services. Unfortunately the service provides no offical API.

## Unofficial API

This repository provides a Go API for programmatic access to the site.

## Unoffical CLI

Included in this repository is a CLI application using the Go API. It's been wrapped as a Docker container. Run it like this:

`docker run docker.pkg.github.com/abrie/mapatlapi/mapatlcli:latest`

## Technical Details

Internally, the [MapATL website](https://egis.atlantaga.gov/app/home/index.html) takes an address and performs three interesting RPC calls. This API is a lightweight wrapper around those calls.

1. _Geocoder_:
	- Input: An address, ex. "123 Peachtree St."
	- Output: A list of candidate locations, each with a `Ref_ID` number.
2. _Location_:
	- Input: A `Ref_ID` number
	- Output: Information about the location. See [this definition](https://github.com/abrie/mapatlapi/blob/master/src/github.com/abrie/mapatlapi/internal/point/response.go) for details.
3. _Places_ (_of interest_):
	- Input: A `Ref_ID` and a `Category`, which can be one of two values: `"PL_PARKS"` or `"TRANS_MARTA_RAIL_STATIONS"`.
	- Output: A list of places matching the category, each having a distance attribute.
