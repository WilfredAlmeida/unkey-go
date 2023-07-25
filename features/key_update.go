package features

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WilfredAlmeida/unkey-go/utils"
)

type NullableField[T any] struct {
	Defined bool
	Value   *T
}

func (m *NullableField[T]) UnmarshalJSON(data []byte) error {
	m.Defined = true
	return json.Unmarshal(data, &m.Value)
}

type RateLimitSchema struct {
	Type           string `json:"type" validate:"required"`
	Limit          int64  `json:"limit" validate:"required"`
	RefillRate     int64  `json:"refillRate" validate:"required"`
	RefillInterval int64  `json:"refillInterval" validate:"required"`
}

type KeyUpdateRequest struct {
	KeyId     string                         `json:"keyId" validate:"required"`
	Name      NullableField[string]          `json:"name"`
	OwnerId   NullableField[string]          `json:"ownerId"`
	Meta      NullableField[map[string]any]  `json:"meta"`
	Expires   NullableField[int64]           `json:"expires"`
	Ratelimit NullableField[RateLimitSchema] `json:"ratelimit"`
	Remaining NullableField[int64]           `json:"remaining"`
}

type keyUpdateJsonRequest struct {
	KeyId     string          `json:"keyId" validate:"required"`
	Name      string          `json:"name,omitempty"`
	OwnerId   string          `json:"ownerId,omitempty"`
	Meta      map[string]any  `json:"meta,omitempty"`
	Expires   int64           `json:"expires,omitempty"`
	Ratelimit *RateLimitSchema `json:"ratelimit,omitempty"`
	Remaining int64           `json:"remaining,omitempty"`
}

type KeyUpdateResponse struct{}

func KeyUpdate(key string, userProvidedData KeyUpdateRequest, authToken string) (KeyUpdateResponse, error) {


	// The API expects values which are to be updated. If you don't want to update a value, don't send it in the request body.
	// Now to determine this in our code, we use the NullableField struct. If the value is not defined, we don't send it in the request body.
	// If the value is defined, we send it in the request body.
	// Do you find this complex? So do I. Blame Golang & check this out https://www.youtube.com/watch?v=xvFZjo5PgG0

	var actualUpdateRequest keyUpdateJsonRequest

	actualUpdateRequest.KeyId = userProvidedData.KeyId

	// The following checks determine if a field is set by user. If yes then it sets it in the api request body
	if userProvidedData.Name.Defined {
		if userProvidedData.Name.Value != nil {
			actualUpdateRequest.Name = *userProvidedData.Name.Value
		} 
	}

	if userProvidedData.OwnerId.Defined {
		if userProvidedData.OwnerId.Value != nil {
			actualUpdateRequest.OwnerId = *userProvidedData.OwnerId.Value
		} 
	}
	if userProvidedData.Meta.Defined {
		if userProvidedData.Meta.Value != nil {
			actualUpdateRequest.Meta = *userProvidedData.Meta.Value
		} 
	}
	if userProvidedData.Expires.Defined {
		if userProvidedData.Expires.Value != nil {
			actualUpdateRequest.Expires = *userProvidedData.Expires.Value
		} 
	}
	if userProvidedData.Ratelimit.Defined {
		if userProvidedData.Ratelimit.Value != nil {
			actualUpdateRequest.Ratelimit = &RateLimitSchema{
				Type:           userProvidedData.Ratelimit.Value.Type,
				Limit:          userProvidedData.Ratelimit.Value.Limit,
				RefillRate:     userProvidedData.Ratelimit.Value.RefillRate,
				RefillInterval: userProvidedData.Ratelimit.Value.RefillInterval,
			}
		}
	}
	if userProvidedData.Remaining.Defined {
		if userProvidedData.Remaining.Value != nil {
			actualUpdateRequest.Remaining = *userProvidedData.Remaining.Value
		}
	}

	// Convert the request body to JSON
	jsonData, err := json.Marshal(actualUpdateRequest)
	if err != nil {
		return KeyUpdateResponse{}, err
	}

	fmt.Println(string(jsonData))

	// Create the HTTP request
	request, err := http.NewRequest("PUT", utils.UNKEY_API_URL+"/keys/"+key, bytes.NewBuffer(jsonData))
	if err != nil {
		return KeyUpdateResponse{}, err
	}

	// Set the request headers
	request.Header.Set("Authorization", "Bearer "+authToken)
	request.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return KeyUpdateResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return KeyUpdateResponse{}, err
	}

	// Handle successful response
	if resp.StatusCode == http.StatusOK {
		// Response is successful, parse the response body as CreateKeyResponse
		var updateKeyResponse KeyUpdateResponse
		err = json.Unmarshal(body, &updateKeyResponse)
		if err != nil {
			return KeyUpdateResponse{}, err
		}

		return updateKeyResponse, err
	} else {
		// Handle error response
		return KeyUpdateResponse{}, fmt.Errorf(string(body))
	}
}
