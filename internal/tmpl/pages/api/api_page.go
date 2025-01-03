package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"

	general "github.com/ropfoo/gothulhu/internal/tmpl/general"
	navigation "github.com/ropfoo/gothulhu/internal/tmpl/navigation"
)

func ApiPage() string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "api", "api_page.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	data := struct {
		Head       template.HTML
		Navigation template.HTML
		Footer     template.HTML
	}{
		Head: general.Head(general.HeadParams{
			Title:       "API | Gothulhu",
			Description: "API for Gothulhu.",
		}),
		Navigation: template.HTML(navigation.Navigation()),
		Footer:     template.HTML(general.Footer()),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return ""
	}

	return buf.String()
}
