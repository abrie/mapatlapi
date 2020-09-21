package location

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseHttpResponse_ErrorStatusCode(t *testing.T) {
	httpResponse := http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
		StatusCode: http.StatusNotFound}

	_, err := ParseHttpResponse(&httpResponse)

	if err == nil {
		t.Fatalf("Expected an error, but got none.")
	}
}

func TestParseHttpResponse_BodyNotJSON(t *testing.T) {
	httpResponse := http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("this-is-not-json"))),
		StatusCode: http.StatusOK}

	if _, err := ParseHttpResponse(&httpResponse); err == nil {
		t.Fatalf("Expected an error, but got none.")
	}
}

func TestParseHttpResponse_ParsesBody(t *testing.T) {
	bodyBytes, err := ioutil.ReadFile("testdata/location.json")
	if err != nil {
		t.Fatalf("Unable to read testdata file: %s", err)
	}

	httpResponse := http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer(bodyBytes)),
		StatusCode: http.StatusOK}

	if _, err := ParseHttpResponse(&httpResponse); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}
