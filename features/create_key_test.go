package features

import (
	"errors"
	"os"
	"testing"
	"reflect"

	"github.com/joho/godotenv"
)

func TestCreateKey(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
	  t.Errorf("Error loading .env file")
	}

	testCases := []struct {
		name           string
		requestBody    CreateKeyRequest
		authToken      string
		expectedResult CreateKeyResponse
		expectedError  error
	}{
		{
			name: "Successful Response",
			requestBody: CreateKeyRequest{
				APIId:      os.Getenv("API_ID"),
				Prefix:     "testPrefix",
				ByteLength: 16,
				OwnerId:    "testOwnerId",
				Meta:       Meta{"key": "value"},
				Expires:    1696951966471,
				Remaining:  10,
				RateLimit: CreateKeyRateLimit{
					Type:           "testType",
					Limit:          100,
					RefillRate:     10,
					RefillInterval: 60,
				},
			},
			authToken: os.Getenv("AUTH_TOKEN"),
			expectedResult: CreateKeyResponse{},
			expectedError: nil,
		},
		{
			name: "Error Response",
			requestBody: CreateKeyRequest{
				APIId:      os.Getenv("API_ID"),
				Prefix:     "testPrefix",
				ByteLength: 16,
				OwnerId:    "testOwnerId",
				Meta:       Meta{"key": "value"},
				Expires:    123456,
				Remaining:  10,
				RateLimit: CreateKeyRateLimit{
					Type:           "testType",
					Limit:          100,
					RefillRate:     10,
					RefillInterval: 60,
				},
			},
			authToken: os.Getenv("AUTH_TOKEN"),
			expectedResult: CreateKeyResponse{},
			expectedError:  errors.New(`{"error":"'expires' must be in the future, did you pass in a timestamp in seconds instead of milliseconds?","code":"BAD_REQUEST"}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			response, err := CreateKey(tc.requestBody, tc.authToken)

			if err != nil && tc.expectedError == nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if err == nil && tc.expectedError != nil {
				t.Errorf("Expected error: %v, got nil", tc.expectedError)
				return
			}
			if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				return
			}

			// Since the returned key is a random string, we can only check the type
			if reflect.TypeOf(response.Key).Kind() != reflect.String {
				t.Errorf("Expected Key: %s, got: %s", tc.expectedResult.Key, response.Key)
			}
			if reflect.TypeOf(response.KeyId).Kind() != reflect.String {
				t.Errorf("Expected KeyId: %s, got: %s", tc.expectedResult.KeyId, response.KeyId)
			}
		})
	}
}
