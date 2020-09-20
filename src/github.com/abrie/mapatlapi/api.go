package mapatlapi

import (
	"context"
	"fmt"
)

import (
	"github.com/abrie/mapatlapi/internal/geocoder"
	"github.com/abrie/mapatlapi/internal/location"
	"github.com/abrie/mapatlapi/internal/places"
	"github.com/abrie/mapatlapi/internal/submitter"
)

func SearchByAddress(ctx context.Context, address string, maxLocations int64) (*geocoder.Response, error) {
	service := &geocoder.Service{
		Endpoint: "https://egis.atlantaga.gov/arc/rest/services/WebLocators/TrAddrPointS/GeocodeServer/findAddressCandidates"}

	request := geocoder.Request{Address: address, MaxLocations: maxLocations}
	httpRequest, err := request.BuildHttpRequest(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to build geocoder request: %w", err)
	}

	httpResponse, err := submitter.Submit(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Geocoder submitter failed: %w", err)
	}

	return geocoder.ParseHttpResponse(httpResponse)
}

func FetchLocation(ctx context.Context, refId int) (*location.Response, error) {
	service := &location.Service{
		Endpoint: "http://egis.atlantaga.gov/app/home/php/egisws.php"}

	request := location.Request{Ref_ID: refId}
	httpRequest, err := request.BuildHttpRequest(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to build location request: %w", err)
	}

	httpResponse, err := submitter.Submit(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Location submitter failed: %w", err)
	}

	return location.ParseHttpResponse(httpResponse)
}

func FetchPlaces(ctx context.Context, refId int, category string) (*places.Response, error) {
	service := &places.Service{
		Endpoint: "http://egis.atlantaga.gov/app/home/php/egispoi.php"}

	request := places.Request{Ref_ID: refId, Category: category}
	httpRequest, err := request.BuildHttpRequest(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to build places request: %w", err)
	}

	httpResponse, err := submitter.Submit(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Places submitter failed: %w", err)
	}

	return places.ParseHttpResponse(httpResponse)
}
