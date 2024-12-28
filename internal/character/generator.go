package character

import "math/rand"

// generate a character with random stats
func GenerateCharacter(name string, age int) Character {
	return Character{
		Name: name,
		Age:  age,
		STR:  rand.Intn(100),
		DEX:  rand.Intn(100),
		CON:  rand.Intn(100),
		INT:  rand.Intn(100),
		WIS:  rand.Intn(100),
		CHA:  rand.Intn(100),
	}
}
