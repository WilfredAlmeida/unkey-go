package features

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type APIGetResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func APIGet(apiID, authToken string) (APIGetResponse, error) {

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", utils.UNKEY_API_URL+"/apis/"+apiID, nil)
	if err != nil {
		return APIGetResponse{}, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", authToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return APIGetResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIGetResponse{}, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return APIGetResponse{}, fmt.Errorf(string(body))
	}

	// Parse the response JSON
	var response APIGetResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return APIGetResponse{}, err
	}

	return response, nil
}
