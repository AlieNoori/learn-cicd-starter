package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name           string
		passHeader     http.Header
		expectedAPIKey string
		wantErr        bool
	}{
		{
			name: "correct header",
			passHeader: http.Header{
				"Authorization": []string{"ApiKey 628a07b885bf2dde8c25262e7c198d8420bf02776299061352e860d0d143d143"},
			},
			expectedAPIKey: "628a07b885bf2dde8c25262e7c198d8420bf02776299061352e860d0d143d143",
			wantErr:        false,
		},
		{
			name: "malformed header",
			passHeader: http.Header{
				"Authorization": []string{"628a07b885bf2dde8c25262e7c198d8420bf02776299061352e860d0d143d143"},
			},
			expectedAPIKey: "",
			wantErr:        true,
		},
		{
			name: "no authorization header",
			passHeader: http.Header{
				"Authorization": []string{""},
			},
			expectedAPIKey: "",
			wantErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualAPIKey, err := GetAPIKey(tc.passHeader)
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("Expected nil - Got %s", err)
				}
				return
			}

			if actualAPIKey != tc.expectedAPIKey {
				t.Fatalf("Expected %s - Got %s", tc.expectedAPIKey, actualAPIKey)
			}
		})
	}
}
