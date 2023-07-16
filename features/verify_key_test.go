package features

import (
	"testing"
)

func TestVerifyKey(t *testing.T) {
	testCases := []struct {
		name           string
		apiKey         string
		expectedResult VerifyKeyResponse
		expectedError   error
	}{
		{
			name:   "True Response",
			apiKey: "key_3ZZSawUTYL1DdsgCycdp7Xdu",
			expectedResult: VerifyKeyResponse{
				Valid: true,
			},
			expectedError: nil,
		},
		{
			name:   "False Response",
			apiKey: "invalidKey_blah_blah",
			expectedResult: VerifyKeyResponse{
				Valid: false,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := VerifyKey(tc.apiKey)

			if err != nil && tc.expectedError == nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if response.Valid != tc.expectedResult.Valid {
				t.Errorf("Expected Valid: %v, got: %v", tc.expectedResult.Valid, response.Valid)
				return
			}
		})
	}
}
