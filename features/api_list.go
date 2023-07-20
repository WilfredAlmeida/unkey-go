package features

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type APIListResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func APIList(apiID, authToken string) (APIListResponse, error) {

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", utils.UNKEY_API_URL+"/apis/"+apiID, nil)
	if err != nil {
		return APIListResponse{}, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", authToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return APIListResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIListResponse{}, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return APIListResponse{}, fmt.Errorf(string(body))
	}

	// Parse the response JSON
	var response APIListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return APIListResponse{}, err
	}

	return response, nil
}
