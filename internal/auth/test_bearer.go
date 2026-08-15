package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name    string
		auth    string
		want    string
		wantErr error
	}{
		{
			name:    "valid bearer token",
			auth:    "Bearer TOKEN_STRING",
			want:    "TOKEN_STRING",
			wantErr: nil,
		},
		{
			name:    "valid lowercase bearer token",
			auth:    "bearer TOKEN_STRING",
			want:    "TOKEN_STRING",
			wantErr: nil,
		},
		{
			name:    "missing authorization header",
			auth:    "",
			want:    "",
			wantErr: ErrNoBearerToken,
		},
		{
			name:    "invalid authorization scheme",
			auth:    "Basic TOKEN_STRING",
			want:    "",
			wantErr: ErrNoBearerToken,
		},
		{
			name:    "malformed authorization header",
			auth:    "Bearer",
			want:    "",
			wantErr: ErrNoBearerToken,
		},
		{
			name:    "empty bearer token",
			auth:    "Bearer ",
			want:    "",
			wantErr: ErrNoBearerToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.auth != "" {
				headers.Set("Authorization", tt.auth)
			}

			got, err := GetBearerToken(headers)
			if err != tt.wantErr {
				t.Fatalf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("GetBearerToken() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		auth    string
		want    string
		wantErr error
	}{
		{
			name:    "valid api key",
			auth:    "ApiKey THE_KEY_HERE",
			want:    "THE_KEY_HERE",
			wantErr: nil,
		},
		{
			name:    "valid lowercase api key",
			auth:    "apikey THE_KEY_HERE",
			want:    "THE_KEY_HERE",
			wantErr: nil,
		},
		{
			name:    "missing authorization header",
			auth:    "",
			want:    "",
			wantErr: ErrNoAPIKey,
		},
		{
			name:    "invalid authorization scheme",
			auth:    "Bearer THE_KEY_HERE",
			want:    "",
			wantErr: ErrNoAPIKey,
		},
		{
			name:    "malformed authorization header",
			auth:    "ApiKey",
			want:    "",
			wantErr: ErrNoAPIKey,
		},
		{
			name:    "empty api key",
			auth:    "ApiKey ",
			want:    "",
			wantErr: ErrNoAPIKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.auth != "" {
				headers.Set("Authorization", tt.auth)
			}

			got, err := GetAPIKey(headers)
			if err != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
