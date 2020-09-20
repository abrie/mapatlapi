package places

import (
	"bytes"
	"context"
	"testing"
)

func TestBuildHttpRequest(t *testing.T) {
	request := Request{Ref_ID: 1001, Category: "PL_PARKS"}
	service := &Service{Endpoint: "https://local.dev"}

	httpRequest, err := request.BuildHttpRequest(context.Background(), service)
	if err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	expectedUrl := service.Endpoint
	gotUrl := httpRequest.URL.String()

	if gotUrl != expectedUrl {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expectedUrl, gotUrl)
	}

	expectedBody := "wsparam%5B%5D=1001&wsparam%5B%5D=PL_PARKS&wsparam%5B%5D=2231146.3740618555&wsparam%5B%5D=1377286.8144072623"
	gotBody := new(bytes.Buffer)
	if _, err := gotBody.ReadFrom(httpRequest.Body); err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	if gotBody.String() != expectedBody {
		t.Errorf("\nWant:\t%s'\nGot:\t'%s'", expectedBody, gotBody.String())
	}
}
