package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
)

// return html template for character
func characterForm(c model.Character) string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_form.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		log.Printf("Error executing template: %v", err)
		return ""
	}

	return buf.String()
}
