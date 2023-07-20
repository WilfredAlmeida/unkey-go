package features

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestAPIGet(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		t.Errorf("Error loading .env file")
	}

	testCases := []struct {
		name           string
		apiId          string
		authToken      string
		expectedResult APIGetResponse
		expectedError  error
	}{
		{
			name:      "Successful Response",
			apiId:     os.Getenv("API_ID"),
			authToken: os.Getenv("AUTH_TOKEN"),
			expectedResult: APIGetResponse{
				ID:          os.Getenv("API_ID"),
				Name:        "first-api",
				WorkspaceID: "ws_C4EkWVE5UFjG4gdZjKJ9wu",
			},
			expectedError: nil,
		},
		{
			name:           "Error Response",
			apiId:          "someId",
			authToken:      os.Getenv("AUTH_TOKEN"),
			expectedResult: APIGetResponse{},
			expectedError:  errors.New(`{"error":"unable to find api: someId","code":"NOT_FOUND"}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			response, err := APIGet(tc.apiId, tc.authToken)

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

			if !reflect.DeepEqual(response, tc.expectedResult) {
				t.Errorf("Expected response: %v, got: %v", tc.expectedResult, response)
				return
			}

		})
	}
}
