package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
)

func GeneratePage(c model.Character) string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "generate_page.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	data := struct {
		CharacterForm   template.HTML
		CharacterResult template.HTML
	}{
		CharacterForm:   template.HTML(characterForm(c)),
		CharacterResult: template.HTML(characterResult(c)),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return ""
	}

	return buf.String()
}
