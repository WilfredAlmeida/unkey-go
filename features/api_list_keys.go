package features

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)


type Key struct {
    ID          string    `json:"id,omitempty"`
    APIID       string    `json:"apiId,omitempty"`
    WorkspaceID string    `json:"workspaceId,omitempty"`
    Start       string    `json:"start,omitempty"`
    CreatedAt   int64     `json:"createdAt,omitempty"`
    Expires     *int64    `json:"expires,omitempty"`
    Ratelimit   ratelimitResponse `json:"ratelimit,omitempty"`
}

type APIListKeysResponse struct {
	Keys  []Key `json:"keys,omitempty"`
	Total int   `json:"total,omitempty"`
}

func APIListKeys(apiID, authToken string) (APIListKeysResponse, error) {

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", utils.UNKEY_API_URL+"/apis/"+apiID+"/keys", nil)
	if err != nil {
		return APIListKeysResponse{}, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", authToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return APIListKeysResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIListKeysResponse{}, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return APIListKeysResponse{}, fmt.Errorf(string(body))
	}

	// Parse the response JSON
	var response APIListKeysResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return APIListKeysResponse{}, err
	}

	return response, nil
}
