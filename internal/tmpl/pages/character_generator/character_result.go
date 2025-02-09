package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"

	"github.com/ropfoo/gothulhu/internal/model"
	"github.com/ropfoo/gothulhu/internal/web"
)

func characterResult(c model.Character) string {
	if c.Name == "" {
		log.Printf("Character name is empty")
		return ""
	}

	tmpl, err := template.ParseFiles(filepath.Join("internal", "tmpl", "pages", "character_generator", "character_result.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return ""
	}

	characterURL := buildCharacterURL(c)
	characterURL = web.ReplaceWhitespacesWithUnderscores(characterURL)

	characterStatRows := ""
	stats := []string{"STR", "DEX", "CON", "INT", "WIS", "CHA"}
	for _, stat := range stats {
		var statValue []float32
		switch stat {
		case "STR":
			statValue = c.Stats.STR
		case "DEX":
			statValue = c.Stats.DEX
		case "CON":
			statValue = c.Stats.CON
		case "INT":
			statValue = c.Stats.INT
		case "WIS":
			statValue = c.Stats.WIS
		case "CHA":
			statValue = c.Stats.CHA
		case "SIZ":
			statValue = c.Stats.SIZ
		}

		characterStatRows += characterStatRow(stat, statValue[0], statValue[1], statValue[2])
	}

	data := struct {
		Character         model.Character
		CharacterURL      string
		CharacterStatRows template.HTML
	}{
		Character:         c,
		CharacterURL:      characterURL,
		CharacterStatRows: template.HTML(characterStatRows),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error executing template: %v", err)
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
		{"str", c.Stats.STR[0]},
		{"dex", c.Stats.DEX[0]},
		{"con", c.Stats.CON[0]},
		{"intl", c.Stats.INT[0]},
		{"wis", c.Stats.WIS[0]},
		{"cha", c.Stats.CHA[0]},
		{"siz", c.Stats.SIZ[0]},
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
