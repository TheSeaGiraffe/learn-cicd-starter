package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_validation(t *testing.T) {
	tests := []struct {
		name        string
		headerKey   string
		headerValue string
		want        string
	}{
		{"no auth header", "Poodonkis", "Bearer 00000", ErrNoAuthHeaderIncluded.Error()},
		{"malformed auth header", "Authorization", "Bearer 00000", "malformed authorization header"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqHeader := http.Header{}
			reqHeader.Add(tt.headerKey, tt.headerValue)

			_, err := GetAPIKey(reqHeader)
			if err != nil {
				if err.Error() != tt.want {
					t.Errorf("GetAPIKey() err == %s; want == %s", err.Error(), tt.want)
				}
			} else {
				t.Error("Expected an error")
			}
		})
	}
}

func TestGetAPIKey_checkWorking(t *testing.T) {
	tests := []struct {
		name        string
		headerKey   string
		headerValue string
		want        string
	}{
		{"valid header with key", "Authorization", "ApiKey 00000", "00000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqHeader := http.Header{}
			reqHeader.Add(tt.headerKey, tt.headerValue)

			apiKey, err := GetAPIKey(reqHeader)
			if err != nil {
				t.Fatalf("GetAPIKey() error: %v", err)
			}

			if apiKey != tt.want {
				t.Errorf("got %q; want %q", apiKey, tt.want)
			}
		})
	}
}
