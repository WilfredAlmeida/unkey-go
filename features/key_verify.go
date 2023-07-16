package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ratelimitResponse struct {
	Limit     int64 `json:"limit"`
	Remaining int64 `json:"remaining"`
	Reset     int64 `json:"reset"`
}

type VerifyKeyResponse struct {
	Valid     bool               `json:"valid"`
	OwnerId   string             `json:"ownerId,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Expires   int64              `json:"expires,omitempty"`
	Remaining int64              `json:"remaining,omitempty"`
	Ratelimit *ratelimitResponse `json:"ratelimit,omitempty"`
	Code      string             `json:"code,omitempty"`
}

type VerifyKeyErrorResponse struct {
	Valid     bool               `json:"valid"`
	Ratelimit *ratelimitResponse `json:"ratelimit,omitempty"`
}

func VerifyKey(apiKey string) (VerifyKeyResponse, error) {
	url := "https://api.unkey.dev/v1/keys/verify"
	data := map[string]string{
		"key": apiKey,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return VerifyKeyResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return VerifyKeyResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return VerifyKeyResponse{}, err
	}
	defer resp.Body.Close()

	var verifyKeyResponse VerifyKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&verifyKeyResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return VerifyKeyResponse{}, err
	}

	// fmt.Println("Response:", verifyKeyResponse)

	return verifyKeyResponse, nil
}
