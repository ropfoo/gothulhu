package character

import (
	"github.com/ropfoo/gothulhu/internal/model"
	"github.com/ropfoo/gothulhu/internal/web"
)

// generate a character with random stats
func GenerateCharacter(params model.CharacterParams) model.Character {
	var age int
	if params.Age == 0 {
		age = getRandomAge()
	} else {
		age = params.Age
	}

	var gender model.Gender
	if params.Gender == "" {
		gender = getRandomGender()
	} else {
		if params.Gender == "male" || params.Gender == "female" {
			gender = model.Gender(params.Gender)
		} else {
			gender = getRandomGender()
		}
	}

	var name string
	if params.Name == "" {
		name = getRandomName(gender)
	} else {
		name = params.Name
		name = web.ReplaceUnderscoresWithWhitespaces(name)
	}

	character := model.Character{
		Name:   name,
		Age:    age,
		Gender: gender,
		Stats: model.Stats{
			STR: getStat(params.Stats.STR[0]),
			DEX: getStat(params.Stats.DEX[0]),
			CON: getStat(params.Stats.CON[0]),
			INT: getStat(params.Stats.INT[0]),
			WIS: getStat(params.Stats.WIS[0]),
			CHA: getStat(params.Stats.CHA[0]),
			SIZ: getStat(params.Stats.SIZ[0]),
		},
	}
	character.HP = getHP(character.Stats.CON[0], character.Stats.SIZ[0])
	character.DamageBonus = getDamageBonus(character.Stats.STR[0], character.Stats.SIZ[0])
	return character
}
