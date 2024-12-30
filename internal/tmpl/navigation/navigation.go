package tmpl

import (
	"bytes"
	"html/template"
)

func Navigation() template.HTML {
	tmpl, err := template.ParseFiles("internal/tmpl/navigation/navigation.html")
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return ""
	}

	return template.HTML(buf.String())
}
