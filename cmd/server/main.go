package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ropfoo/gothulhu/internal/character"
	"github.com/ropfoo/gothulhu/internal/handlers"
	"github.com/ropfoo/gothulhu/internal/web"
)

func main() {
	// create an endpoint to generate a character
	http.HandleFunc("/api/generate", func(w http.ResponseWriter, r *http.Request) {
		characterParams, err := web.GetCharacterParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		character := character.GenerateCharacter(characterParams)

		log.Printf("Generating character for %s", characterParams.Name)
		log.Printf("Character: %+v", character)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(character.ToJSON()))
	})

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

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)
	http.HandleFunc("/api-docs", handlers.ApiDocsHandler)

	const port = 8000
	log.Printf("Starting server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
