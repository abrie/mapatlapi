package location

import (
	"bytes"
	"context"
	"testing"
)

func TestBuildHttpRequest(t *testing.T) {
	request := Request{Ref_ID: 1001}
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

	expectedBody := "wsparam%5B%5D=1001"
	gotBody := new(bytes.Buffer)
	if err := gotBody.ReadFrom(httpRequest.Body); err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	if gotBody.String() != expectedBody {
		t.Errorf("\nWant:\t%s'\nGot:\t'%s'", expectedBody, gotBody.String())
	}
}
