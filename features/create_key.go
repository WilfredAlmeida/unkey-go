package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type CreateKeyRequest struct {
	APIId      string             `json:"apiId"`
	Prefix     string             `json:"prefix"`
	ByteLength int                `json:"byteLength"`
	OwnerId    string             `json:"ownerId"`
	Meta       Meta               `json:"meta"`
	Expires    int64              `json:"expires"`
	Remaining  int                `json:"remaining"`
	RateLimit  CreateKeyRateLimit `json:"ratelimit"`
}

type Meta map[string]string

type CreateKeyRateLimit struct {
	Type           string `json:"type"`
	Limit          int    `json:"limit"`
	RefillRate     int    `json:"refillRate"`
	RefillInterval int    `json:"refillInterval"`
}

type CreateKeyResponse struct {
	Key   string `json:"key"`
	KeyId string `json:"keyId"`
}

func CreateKey(requestBody CreateKeyRequest, authToken string) (CreateKeyResponse, error) {
	// Convert the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return CreateKeyResponse{}, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", utils.UNKEY_API_URL+"/keys", bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateKeyResponse{}, err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CreateKeyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateKeyResponse{}, err
	}

	// Handle successful response
	if resp.StatusCode == http.StatusOK {
		// Response is successful, parse the response body as CreateKeyResponse
		var createKeyResponse CreateKeyResponse
		err = json.Unmarshal(body, &createKeyResponse)
		if err != nil {
			return CreateKeyResponse{}, err
		}

		return createKeyResponse, err
	} else {
		// Handle error response
		return CreateKeyResponse{}, fmt.Errorf(string(body))
	}
}
