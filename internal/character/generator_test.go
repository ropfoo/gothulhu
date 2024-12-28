package character

import (
	"testing"
)

func TestGenerateCharacter(t *testing.T) {
	// Test case inputs
	name := "Test Character"
	age := 25

	// Generate character
	character := GenerateCharacter(name, age)

	// Test name and age are set correctly
	if character.Name != name {
		t.Errorf("Expected name to be %s, got %s", name, character.Name)
	}
	if character.Age != age {
		t.Errorf("Expected age to be %d, got %d", age, character.Age)
	}

	// Test that stats are within expected range (assuming getRandomStat returns 0-100)
	stats := []int{
		character.STR,
		character.DEX,
		character.CON,
		character.INT,
		character.WIS,
		character.CHA,
		character.SIZ,
	}

	for _, stat := range stats {
		if stat < 0 || stat > 100 {
			t.Errorf("Stat value %d is outside expected range of 0-100", stat)
		}
	}

	// Test that HP and DamageBonus are calculated
	expectedHP := getHP(character.CON, character.SIZ)
	if character.HP != expectedHP {
		t.Errorf("Expected HP to be %d, got %d", expectedHP, character.HP)
	}

	expectedDamageBonus := getDamageBonus(character.STR, character.SIZ)
	if character.DamageBonus != expectedDamageBonus {
		t.Errorf("Expected DamageBonus to be %s, got %s", expectedDamageBonus, character.DamageBonus)
	}
}
