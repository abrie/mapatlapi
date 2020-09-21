package places

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseHttpResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response not OK: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Failed unmarshal response body: %v", err)
	}

	return &response, nil
}
