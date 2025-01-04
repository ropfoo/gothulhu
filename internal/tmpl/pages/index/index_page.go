package tmpl

import (
	"bytes"
	"html/template"
	"path/filepath"

	general "github.com/ropfoo/gothulhu/internal/tmpl/general"
	navigation "github.com/ropfoo/gothulhu/internal/tmpl/navigation"
)

func IndexPage() string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "index", "index_page.html"))
	if err != nil {
		return ""
	}

	data := struct {
		Head       template.HTML
		Navigation template.HTML
		Footer     template.HTML
	}{
		Head: general.Head(general.HeadParams{
			Title:       "Gothulhu",
			Description: "Generate Call of Cthulhu 7th edition characters quickly and easily with Gothulhu - a free character generator tool focused on NPCs.",
		}),
		Navigation: template.HTML(navigation.Navigation()),
		Footer:     template.HTML(general.Footer()),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return ""
	}

	return buf.String()
}
