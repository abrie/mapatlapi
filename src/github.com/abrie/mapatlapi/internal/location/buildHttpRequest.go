package location

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (params *Request) BuildHttpRequest(ctx context.Context, service *Service) (*http.Request, error) {
	data := url.Values{}
	data.Add("wsparam[]", strconv.Itoa(params.Ref_ID))
	encodedData := data.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", service.Endpoint, strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))

	return req, nil
}
