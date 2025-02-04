package character

import (
	"testing"

	"github.com/ropfoo/gothulhu/internal/model"
)

func TestGenerateCharacter(t *testing.T) {
	// Test case inputs
	name := "Test Character"
	age := 25

	// Generate character
	character := GenerateCharacter(model.CharacterParams{
		Name: name,
		Age:  age,
	})

	// Test name and age are set correctly
	if character.Name != name {
		t.Errorf("Expected name to be %s, got %s", name, character.Name)
	}
	if character.Age != age {
		t.Errorf("Expected age to be %d, got %d", age, character.Age)
	}

	// Test that stats are within expected range (assuming getRandomStat returns 0-100)
	stats := []float32{
		character.Stats.STR[0],
		character.Stats.DEX[0],
		character.Stats.CON[0],
		character.Stats.INT[0],
		character.Stats.WIS[0],
		character.Stats.CHA[0],
		character.Stats.SIZ[0],
	}

	for _, stat := range stats {
		if stat < 0 || stat > 100 {
			t.Errorf("Stat value %d is outside expected range of 0-100", stat)
		}
	}

	// Test that HP and DamageBonus are calculated
	expectedHP := getHP(character.Stats.CON[0], character.Stats.SIZ[0])
	if character.HP != expectedHP {
		t.Errorf("Expected HP to be %d, got %d", expectedHP, character.HP)
	}

	expectedDamageBonus := getDamageBonus(character.Stats.STR[0], character.Stats.SIZ[0])
	if character.DamageBonus != expectedDamageBonus {
		t.Errorf("Expected DamageBonus to be %s, got %s", expectedDamageBonus, character.DamageBonus)
	}
}
