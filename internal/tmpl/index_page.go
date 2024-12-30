package tmpl

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func IndexPage() string {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "index_page.html"))
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return ""
	}

	return buf.String()
}
