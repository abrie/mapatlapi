package places

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/* These two values are returned as x and y by the following call to arcgisonline:
https://tasks.arcgisonline.com/ArcGIS/rest/services/Geometry/GeometryServer/project?f=json&outSR=102667&inSR=4326&geometries=%7B%22geometryType%22%3A%22esriGeometryPoint%22%2C%22geometries%22%3A%5B%7B%22x%22%3A-84.38204681928369%2C%22y%22%3A33.78607792976518%2C%22spatialReference%22%3A%7B%22wkid%22%3A4326%7D%7D%5D%7D

The values appear to be stable, and are thus hardcoded here until required otherwise. If these values are omitted, no results are returned.
*/
const wsparam_x = "2231146.3740618555"
const wsparam_y = "1377286.8144072623"

func (params *Request) BuildHttpRequest(ctx context.Context, service *Service) (*http.Request, error) {
	data := url.Values{}
	data.Add("wsparam[]", strconv.Itoa(params.Ref_ID))
	data.Add("wsparam[]", params.Category)
	data.Add("wsparam[]", wsparam_x)
	data.Add("wsparam[]", wsparam_y)
	encodedData := data.Encode()

	fmt.Sprintf("%s", string(encodedData))
	req, err := http.NewRequestWithContext(ctx, "POST", service.Endpoint, strings.NewReader(encodedData))
	if err != nil {
		return nil, fmt.Errorf("Failed to build HTTP request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))

	return req, nil
}
