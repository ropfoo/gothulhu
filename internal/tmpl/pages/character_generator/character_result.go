package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
	"github.com/ropfoo/gothulhu/internal/web"
)

func characterResult(c model.Character) string {
	if c.Name == "" {
		return ""
	}

	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_result.html"))
	if err != nil {
		return ""
	}

	characterURL := buildCharacterURL(c)
	characterURL = web.ReplaceWhitespacesWithUnderscores(characterURL)

	data := struct {
		Character    model.Character
		CharacterURL string
	}{
		Character:    c,
		CharacterURL: characterURL,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return ""
	}

	return buf.String()
}

func buildCharacterURL(c model.Character) string {
	params := []struct {
		key   string
		value interface{}
	}{
		{"name", c.Name},
		{"age", c.Age},
		{"gender", c.Gender},
		{"str", c.Stats.STR},
		{"dex", c.Stats.DEX},
		{"con", c.Stats.CON},
		{"intl", c.Stats.INT},
		{"wis", c.Stats.WIS},
		{"cha", c.Stats.CHA},
		{"siz", c.Stats.SIZ},
	}

	url := "https://gothulhu.fly.dev/generate?"
	for i, p := range params {
		if i > 0 {
			url += "&"
		}
		url += fmt.Sprintf("%s=%v", p.key, p.value)
	}
	return url
}
