package features

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDeleteKey(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		t.Errorf("Error loading .env file")
	}

	testCases := []struct {
		name              string
		keyID             string
		authToken         string
		expectedResult    map[string]interface{}
		expectedError     error
		expectedErrorCode string
	}{
		{
			name:           "Successful Response",
			keyID:          "key_7UvaP1DTRv3boJHc1Y6Lu4", // replace this else test will fail
			authToken:      os.Getenv("AUTH_TOKEN"),
			expectedResult: map[string]interface{}{},
			expectedError:  nil,
		},
		{
			name:              "Error Response",
			keyID:             "key_XpudYiM8Kz5zCh3TEdBooL",
			authToken:         os.Getenv("AUTH_TOKEN"),
			expectedResult:    map[string]interface{}{},
			expectedError:     errors.New(`{"error":"key key_XpudYiM8Kz5zCh3TEdBooL does not exist","code":"NOT_FOUND"}`),
			expectedErrorCode: "404",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := RevokeKey(tc.keyID, tc.authToken)

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
		})
	}
}
