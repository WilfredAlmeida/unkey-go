package features

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAPIListKeys(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		t.Errorf("Error loading .env file")
	}

	testCases := []struct {
		name           string
		apiId          string
		authToken      string
		expectedResult APIListKeysResponse
		expectedError  error
	}{
		{
			name:           "Successful Response",
			apiId:          os.Getenv("API_ID"),
			authToken:      os.Getenv("AUTH_TOKEN"),
			expectedResult: APIListKeysResponse{}, // It'll list lots of keys which we cannot match here
			expectedError:  nil,
		},
		{
			name:           "Error Response",
			apiId:          "someId",
			authToken:      os.Getenv("AUTH_TOKEN"),
			expectedResult: APIListKeysResponse{},
			expectedError:  errors.New(`{"error":"unable to find api: someId","code":"NOT_FOUND"}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := APIListKeys(tc.apiId, tc.authToken)

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

			// If there's no error then it means we got a successful response

		})
	}
}
