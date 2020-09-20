package geocoder

import (
	"context"
	"testing"
)

func TestBuildHttpRequest(t *testing.T) {
	request := Request{Address: "123 Fake St", MaxLocations: 3}
	service := &Service{Endpoint: "https:/local.dev"}

	httpRequest, err := request.BuildHttpRequest(context.Background(), service)

	if err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	got := httpRequest.URL.String()
	expected := `https:///local.dev?Single+Line+Input=123+Fake+St&f=json&maxLocations=3&outFields=%2A&outSR=%7B%22wkid%22%3A4326%7D`

	if got != expected {
		t.Errorf("\nWant:\t'%s'\nGot:\t'%s'", expected, got)
	}

}
