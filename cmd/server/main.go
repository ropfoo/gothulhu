package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ropfoo/gothulhu/internal/character"
	"github.com/ropfoo/gothulhu/internal/model"
	"github.com/ropfoo/gothulhu/internal/tmpl"

	apiPage "github.com/ropfoo/gothulhu/internal/tmpl/pages/api"
)

func main() {
	// create an endpoint to generate a character
	http.HandleFunc("/api/generate", func(w http.ResponseWriter, r *http.Request) {
		characterParams := getCharacterParams(r)
		character := character.GenerateCharacter(characterParams)

		log.Printf("Generating character for %s", characterParams.Name)
		log.Printf("Character: %+v", character)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(character.ToJSON()))
	})

	http.HandleFunc("/styling/main.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")

		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			http.ServeFile(w, r, "styling/main.css")
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Read and serve the file content through gzip
		content, err := os.ReadFile("styling/main.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		gz.Write(content)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(tmpl.IndexPage()))
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		characterParams := getCharacterParams(r)

		var char model.Character
		if characterParams.Name != "" && characterParams.Age != 0 {
			char = character.GenerateCharacter(characterParams)
		}

		log.Printf("Character: %+v", char)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(tmpl.GeneratePage(char)))
	})

	http.HandleFunc("/api-docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(apiPage.ApiPage()))
	})

	const port = 8000
	log.Printf("Starting server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getCharacterParams(r *http.Request) model.CharacterParams {
	fields := []string{"name", "age", "str", "dex", "con", "siz", "intl", "wis", "cha"}
	values := make(map[string]string)

	for _, field := range fields {
		if r.URL.Query().Get(field) != "" {
			values[field] = r.URL.Query().Get(field)
		} else {
			values[field] = r.FormValue(field)
		}
	}

	ageInt, _ := strconv.Atoi(values["age"])
	strInt, _ := strconv.Atoi(values["str"])
	dexInt, _ := strconv.Atoi(values["dex"])
	conInt, _ := strconv.Atoi(values["con"])
	sizInt, _ := strconv.Atoi(values["siz"])
	intlInt, _ := strconv.Atoi(values["intl"])
	wisInt, _ := strconv.Atoi(values["wis"])
	chaInt, _ := strconv.Atoi(values["cha"])
	return model.CharacterParams{
		Name: values["name"],
		Age:  ageInt,
		Stats: model.Stats{
			STR: strInt,
			DEX: dexInt,
			CON: conInt,
			SIZ: sizInt,
			INT: intlInt,
			WIS: wisInt,
			CHA: chaInt,
		},
	}
}
