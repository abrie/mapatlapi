package geocoder

import (
	"context"
	"net/http"
	"strconv"
)

func (params *Request) BuildHttpRequest(ctx context.Context, service *Service) (*http.Request, error) {
	req, err := http.NewRequestWiteContext(ctx, "GET", service.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("Single Line Input", params.Address)
	q.Add("f", "json")
	q.Add("outFields", "*")
	q.Add("outSR", `{"wkid":4326}`)
	q.Add("maxLocations", strconv.FormatInt(params.MaxLocations, 10))

	req.URL.RawQuery = q.Encode()

	return req, nil
}
