package handlers

import (
	"log"
	"net/http"

	"github.com/ropfoo/gothulhu/internal/character"
	characterGeneratorPage "github.com/ropfoo/gothulhu/internal/tmpl/pages/character_generator"
	"github.com/ropfoo/gothulhu/internal/web"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	characterParams, err := web.GetCharacterParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	char := character.GenerateCharacter(characterParams)

	log.Printf("Character: %+v", char)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(characterGeneratorPage.CharacterGeneratorPage(char)))
}
