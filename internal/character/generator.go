package character

// generate a character with random stats
func GenerateCharacter(name string, age int) Character {
	character := Character{
		Name: name,
		Age:  age,
		STR:  getRandomStat(),
		DEX:  getRandomStat(),
		CON:  getRandomStat(),
		INT:  getRandomStat(),
		WIS:  getRandomStat(),
		CHA:  getRandomStat(),
		SIZ:  getRandomStat(),
	}
	character.HP = getHP(character.CON, character.SIZ)
	character.DamageBonus = getDamageBonus(character.STR, character.SIZ)
	return character
}
