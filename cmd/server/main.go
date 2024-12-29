package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ropfoo/gothulhu/internal/character"
	"github.com/ropfoo/gothulhu/internal/model"
	"github.com/ropfoo/gothulhu/internal/tmpl"
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
		http.ServeFile(w, r, "styling/main.css")
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
