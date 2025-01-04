package web

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeCSSFile(t *testing.T) {
	// Create a temporary CSS file for testing
	tempDir := t.TempDir()
	cssContent := "body { color: blue; }"
	testFile := filepath.Join(tempDir, "test.css")
	if err := os.WriteFile(testFile, []byte(cssContent), 0666); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name           string
		acceptEncoding string
		wantCompressed bool
		wantStatusCode int
	}{
		{
			name:           "with gzip encoding",
			acceptEncoding: "gzip",
			wantCompressed: true,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "without gzip encoding",
			acceptEncoding: "",
			wantCompressed: false,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test request
			req := httptest.NewRequest("GET", "/test.css", nil)
			if tt.acceptEncoding != "" {
				req.Header.Set("Accept-Encoding", tt.acceptEncoding)
			}

			// Create a response recorder
			w := httptest.NewRecorder()

			// Call the function being tested
			ServeCSSFile(w, req, testFile)

			// Check status code
			if w.Code != tt.wantStatusCode {
				t.Errorf("ServeCSSFile() status = %v, want %v", w.Code, tt.wantStatusCode)
			}

			// Check Cache-Control header
			if got := w.Header().Get("Cache-Control"); got != "public, max-age=31536000" {
				t.Errorf("ServeCSSFile() Cache-Control = %v, want public, max-age=31536000", got)
			}

			// Get response body
			resp := w.Result()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check content
			if tt.wantCompressed {
				// For gzipped content, decompress and verify
				if got := resp.Header.Get("Content-Encoding"); got != "gzip" {
					t.Errorf("ServeCSSFile() Content-Encoding = %v, want gzip", got)
				}

				reader, err := gzip.NewReader(w.Body)
				if err != nil {
					t.Fatal(err)
				}
				decompressed, err := io.ReadAll(reader)
				if err != nil {
					t.Fatal(err)
				}
				if string(decompressed) != cssContent {
					t.Errorf("ServeCSSFile() decompressed content = %v, want %v", string(decompressed), cssContent)
				}
			} else {
				// For non-gzipped content, verify directly
				if string(body) != cssContent {
					t.Errorf("ServeCSSFile() content = %v, want %v", string(body), cssContent)
				}
			}
		})
	}
}
