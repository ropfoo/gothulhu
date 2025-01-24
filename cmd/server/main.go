package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ropfoo/gothulhu/internal/handlers"
	"github.com/ropfoo/gothulhu/internal/handlers/api"
	"github.com/ropfoo/gothulhu/internal/web"
)

func main() {

	http.HandleFunc("/styling/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".css") {
			http.NotFound(w, r)
			return
		}
		web.ServeCSSFile(w, r, r.URL.Path[1:]) // Remove leading slash
	})

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		web.ServeAssetFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/favicon.ico")
	})

	// Web
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)
	http.HandleFunc("/api-docs", handlers.ApiDocsHandler)

	// API
	http.HandleFunc("/api/generate", api.GenerateHandler)

	const port = 8000
	log.Printf("Starting server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
