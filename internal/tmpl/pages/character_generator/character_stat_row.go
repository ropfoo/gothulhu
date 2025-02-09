package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
)

func characterStatRow(title string, stat float32, mod float32, mod2 float32) string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_stat_row.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	data := struct {
		Title string
		Stat  float32
		Mod   float32
		Mod2  float32
	}{
		Title: title,
		Stat:  stat,
		Mod:   mod,
		Mod2:  mod2,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return ""
	}

	return buf.String()
}
