package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/WilfredAlmeida/unkey-go/utils"
	"net/http"
	"time"
)

type ratelimitResponse struct {
	Limit     int64 `json:"limit"`
	Remaining int64 `json:"remaining"`
	Reset     int64 `json:"reset"`
}

type KeyVerifyResponse struct {
	Valid       bool                   `json:"valid"`
	OwnerId     string                 `json:"ownerId,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Expires     int64                  `json:"expires,omitempty"`
	Remaining   int64                  `json:"remaining,omitempty"`
	Ratelimit   *ratelimitResponse     `json:"ratelimit,omitempty"`
	Code        string                 `json:"code,omitempty"`
	Enabled     bool                   `json:"enabled"`
	Environment string                 `json:"environment"`
	KeyId       string                 `json:"keyId"`
	Name        string                 `json:"name"`
	Permissions []string               `json:"permissions"`
}

func KeyVerify(apiKey string) (KeyVerifyResponse, error) {

	data := map[string]string{
		"key": apiKey,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return KeyVerifyResponse{}, err
	}

	req, err := http.NewRequest("POST", utils.UNKEY_API_URL+"/keys/verify", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return KeyVerifyResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return KeyVerifyResponse{}, err
	}
	defer resp.Body.Close()

	var verifyKeyResponse KeyVerifyResponse
	err = json.NewDecoder(resp.Body).Decode(&verifyKeyResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return KeyVerifyResponse{}, err
	}

	return verifyKeyResponse, nil
}
