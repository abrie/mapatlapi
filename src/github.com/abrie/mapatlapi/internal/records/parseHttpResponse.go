package records

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseHttpResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response not OK: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %v", err)
	}

	var arr = make([]Record, 0)
	if err := json.Unmarshal(body, &arr); err != nil {
		return nil, fmt.Errorf("Failed unmarshal response body: %v", err)
	}

	return &Response{Records: arr}, nil
}
