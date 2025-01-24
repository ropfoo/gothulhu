package api

import (
	"log"
	"net/http"

	"github.com/ropfoo/gothulhu/internal/character"
	"github.com/ropfoo/gothulhu/internal/web"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
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
}
