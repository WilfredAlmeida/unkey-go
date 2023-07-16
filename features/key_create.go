package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type KeyCreateRequest struct {
	APIId      string             `json:"apiId"`
	Prefix     string             `json:"prefix"`
	ByteLength int                `json:"byteLength"`
	OwnerId    string             `json:"ownerId"`
	Meta       Meta               `json:"meta"`
	Expires    int64              `json:"expires"`
	Remaining  int                `json:"remaining"`
	RateLimit  KeyCreateRateLimit `json:"ratelimit"`
}

type Meta map[string]string

type KeyCreateRateLimit struct {
	Type           string `json:"type"`
	Limit          int    `json:"limit"`
	RefillRate     int    `json:"refillRate"`
	RefillInterval int    `json:"refillInterval"`
}

type KeyCreateResponse struct {
	Key   string `json:"key"`
	KeyId string `json:"keyId"`
}

func KeyCreate(requestBody KeyCreateRequest, authToken string) (KeyCreateResponse, error) {
	// Convert the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return KeyCreateResponse{}, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", utils.UNKEY_API_URL+"/keys", bytes.NewBuffer(jsonData))
	if err != nil {
		return KeyCreateResponse{}, err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return KeyCreateResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return KeyCreateResponse{}, err
	}

	// Handle successful response
	if resp.StatusCode == http.StatusOK {
		// Response is successful, parse the response body as CreateKeyResponse
		var createKeyResponse KeyCreateResponse
		err = json.Unmarshal(body, &createKeyResponse)
		if err != nil {
			return KeyCreateResponse{}, err
		}

		return createKeyResponse, err
	} else {
		// Handle error response
		return KeyCreateResponse{}, fmt.Errorf(string(body))
	}
}
