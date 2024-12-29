package model

import (
	"encoding/json"
	"testing"
)

func TestCharacter_ToJSON(t *testing.T) {
	tests := []struct {
		name     string
		char     Character
		expected string
	}{
		{
			name: "valid character",
			char: Character{
				Name: "Conan",
				Age:  30,
				HP:   75,
				Stats: Stats{
					STR: 85,
					DEX: 70,
					SIZ: 65,
					CON: 80,
					INT: 60,
					WIS: 55,
					CHA: 75,
				},
				DamageBonus: DamageBonus{},
			},
			expected: `{"name":"Conan","age":30,"hp":75,"damage_bonus":{"damage":"","stature":""},"stats":{"str":85,"dex":70,"con":80,"int":60,"wis":55,"cha":75,"siz":65}}`,
		},
		{
			name:     "empty character",
			char:     Character{},
			expected: `{"name":"","age":0,"hp":0,"damage_bonus":{"damage":"","stature":""},"stats":{"str":0,"dex":0,"con":0,"int":0,"wis":0,"cha":0,"siz":0}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.char.ToJSON()

			// Verify the JSON string is valid
			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(result), &parsed); err != nil {
				t.Errorf("ToJSON() produced invalid JSON: %v", err)
			}

			// Compare with expected result
			if result != tt.expected {
				t.Errorf("ToJSON() = %v, want %v", result, tt.expected)
			}
		})
	}
}
