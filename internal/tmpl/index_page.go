package tmpl

import (
	"bytes"
	"html/template"
	"path/filepath"

	navigation "github.com/ropfoo/gothulhu/internal/tmpl/navigation"
)

func IndexPage() string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "index_page.html"))
	if err != nil {
		return ""
	}

	data := struct {
		Navigation template.HTML
	}{
		Navigation: template.HTML(navigation.Navigation()),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return ""
	}

	return buf.String()
}
