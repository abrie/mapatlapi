package mapatlapi

import (
	"context"
	"fmt"
)

import (
	"github.com/abrie/mapatlapi/internal/geocoder"
	"github.com/abrie/mapatlapi/internal/records"
	"github.com/abrie/mapatlapi/internal/submitter"
)

func SearchByAddress(ctx context.Context, address string) (*geocoder.Response, error) {
	service := &geocoder.Service{
		Endpoint: "https://egis.atlantaga.gov/arc/rest/services/WebLocators/TrAddrPointS/GeocodeServer/findAddressCandidates"}

	request := geocoder.Request{Address: address}
	httpRequest, err := request.BuildHttpRequest(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to build geocoder request: %s", err.Error())
	}

	httpResponse, err := submitter.Submit(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Geocoder submitter failed: %s", err.Error())
	}

	return geocoder.ParseHttpResponse(httpResponse)
}

func FetchRecord(ctx context.Context, refId int) (*records.Response, error) {
	service := &records.Service{
		Endpoint: "http://egis.atlantaga.gov/app/home/php/egisws.php"}

	request := records.Request{Ref_ID: refId}
	httpRequest, err := request.BuildHttpRequest(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to build records request: %s", err.Error())
	}

	httpResponse, err := submitter.Submit(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Records submitter failed: %s", err.Error())
	}

	return records.ParseHttpResponse(httpResponse)

}
