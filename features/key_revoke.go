package features

import (
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

func RevokeKey(keyID string, authToken string) (bool, error) {
	// Create the HTTP request
	req, err := http.NewRequest("DELETE", utils.UNKEY_API_URL+"/keys/"+keyID, nil)
	if err != nil {
		return false, err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+authToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, fmt.Errorf(string(body))
	}
}
