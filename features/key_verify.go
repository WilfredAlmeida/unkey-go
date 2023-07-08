package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type VerifyRequestSchema struct {
	Key string `json:"key"`
}

type VerifyResponseSchema struct {
	Valid *bool `json:"valid"`
}

func VerifyRequest(key string) (*bool, error) {
	
	verifyReq := VerifyRequestSchema{
		Key: key,
	}
	payload, err := json.Marshal(verifyReq)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", utils.UNKEY_API_URL+"/keys/verify", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var verifyResp VerifyResponseSchema
	err = json.Unmarshal(body, &verifyResp)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return nil, err
	}

	return verifyResp.Valid, nil
}