package geocoder

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

/* https://egis.atlantaga.gov/arc/rest/services/WebLocators/TrAddrPointS/GeocodeServer/ */
func (params *Request) BuildHttpRequest(ctx context.Context, service *Service) (*http.Request, error) {
	url := strings.Join([]string{service.Endpoint}, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP Request: %v", err)
	}

	q := req.URL.Query()
	q.Add("Single Line Input", params.Address)
	q.Add("f", "json")
	q.Add("outFields", "*")
	q.Add("outSR", `{"wkid":4326}`)
	q.Add("maxLocations", "6")

	req.URL.RawQuery = q.Encode()

	return req, nil
}
