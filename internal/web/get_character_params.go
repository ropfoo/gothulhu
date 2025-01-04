package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ropfoo/gothulhu/internal/model"
)

// GetCharacterParams extracts and validates character parameters from an HTTP request
func GetCharacterParams(r *http.Request) (model.CharacterParams, error) {
	// Extract all form values
	paramValues := extractFormValues(r)

	// Convert and validate numeric stats
	stats, err := parseStats(paramValues)
	if err != nil {
		return model.CharacterParams{}, err
	}

	// Add validation for age
	if _, err := strconv.Atoi(paramValues["age"]); err != nil {
		return model.CharacterParams{}, fmt.Errorf("invalid age: %s", paramValues["age"])
	}

	// Add validation for gender
	gender := paramValues["gender"]
	if gender != "male" && gender != "female" {
		return model.CharacterParams{}, fmt.Errorf("invalid gender: %s", gender)
	}

	return model.CharacterParams{
		Name:   paramValues["name"],
		Gender: model.Gender(paramValues["gender"]),
		Age:    stats["age"],
		Stats:  buildCharacterStats(stats),
	}, nil
}

// extractFormValues gets values from both query parameters and form data
func extractFormValues(r *http.Request) map[string]string {
	fields := []string{"name", "age", "gender", "str", "dex", "con", "siz", "intl", "wis", "cha"}
	values := make(map[string]string)

	for _, field := range fields {
		if v := r.URL.Query().Get(field); v != "" {
			values[field] = v
		} else {
			values[field] = r.FormValue(field)
		}
	}
	return values
}

// parseStats converts string values to integers and returns 0 for invalid values
func parseStats(values map[string]string) (map[string]int, error) {
	statFields := []string{"age", "str", "dex", "con", "siz", "intl", "wis", "cha"}
	stats := make(map[string]int)

	for _, field := range statFields {
		val, err := strconv.Atoi(values[field])
		if err != nil {
			stats[field] = 0 // Set to 0 instead of returning error
			continue
		}
		stats[field] = val
	}
	return stats, nil
}

// buildCharacterStats creates a Stats struct from the parsed values
func buildCharacterStats(stats map[string]int) model.Stats {
	return model.Stats{
		STR: stats["str"],
		DEX: stats["dex"],
		CON: stats["con"],
		SIZ: stats["siz"],
		INT: stats["intl"],
		WIS: stats["wis"],
		CHA: stats["cha"],
	}
}
