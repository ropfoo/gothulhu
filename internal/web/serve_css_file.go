package web

import (
	"compress/gzip"
	"net/http"
	"os"
	"strings"
)

func ServeCSSFile(w http.ResponseWriter, r *http.Request, filepath string) {
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	// Check if the client accepts gzip encoding
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		http.ServeFile(w, r, filepath)
		return
	}

	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()

	// Read and serve the file content through gzip
	content, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	gz.Write(content)
}
