package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"

	general "github.com/ropfoo/gothulhu/internal/tmpl/general"
	navigation "github.com/ropfoo/gothulhu/internal/tmpl/navigation"

	"github.com/ropfoo/gothulhu/internal/model"
)

func CharacterGeneratorPage(c model.Character) string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_generator_page.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	var pageTitle string = "Generate Character | Gothulhu"
	if c.Name != "" {
		pageTitle = c.Name + " | Gothulhu"
	}

	data := struct {
		Head            template.HTML
		Navigation      template.HTML
		CharacterForm   template.HTML
		CharacterResult template.HTML
		Footer          template.HTML
	}{
		Head: general.Head(general.HeadParams{
			Title:       pageTitle,
			Description: "Generate a character for Call of Cthulhu 7th edition with Gothulhu.",
			Stylesheets: []string{"/styling/character.css"},
		}),
		Navigation:      template.HTML(navigation.Navigation()),
		CharacterForm:   template.HTML(characterForm(c)),
		CharacterResult: template.HTML(characterResult(c)),
		Footer:          template.HTML(general.Footer()),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return ""
	}

	return buf.String()
}
