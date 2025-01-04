package tmpl

import (
	"bytes"
	"html/template"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
)

func characterResult(c model.Character) string {
	if c.Name == "" {
		return ""
	}

	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_result.html"))
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return ""
	}

	return buf.String()
}
