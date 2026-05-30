package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Ini adalah pola "Table-Driven Tests"
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-super-secret-key"},
			},
			expectedKey:   "my-super-secret-key",
			expectedError: false,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{}, // Header kosong
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Malformed Header - Missing Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"}, // Tidak ada kunci setelah spasi
			},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Wrong Authorization Scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer my-super-secret-key"}, // Menggunakan Bearer, bukan ApiKey
			},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			// Cek apakah error yang didapat sesuai dengan ekspektasi kita
			if (err != nil) != tc.expectedError {
				t.Fatalf("Test '%s' gagal: expected error %v, got %v", tc.name, tc.expectedError, err)
			}

			// Cek apakah key yang diekstrak sesuai dengan ekspektasi kita
			if key != tc.expectedKey {
				t.Fatalf("Test '%s' gagal: expected key '%v', got '%v'", tc.name, tc.expectedKey, key)
			}
		})
	}
}
