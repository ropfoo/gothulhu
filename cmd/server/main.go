package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ropfoo/gothulhu/internal/character"
)

func main() {
	// create an endpoint to generate a character
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		ageInt, _ := strconv.Atoi(age)
		character := character.GenerateCharacter(name, ageInt)

		log.Printf("Generating character for %s", name)
		log.Printf("Character: %+v", character)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(character.ToJSON()))
	})

	const port = 8000
	log.Printf("Starting server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
