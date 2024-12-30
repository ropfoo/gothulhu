package tmpl

import (
	"bytes"
	"html/template"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
)

// return html template for character
func characterForm(c model.Character) string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "character_form.html"))
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return ""
	}

	return buf.String()
}
