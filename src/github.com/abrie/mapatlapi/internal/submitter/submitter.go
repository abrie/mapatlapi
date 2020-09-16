package submitter

import (
	"net/http"
)

func Submit(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(request)
}
