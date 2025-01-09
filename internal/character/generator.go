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
			STR: getStat(params.Stats.STR),
			DEX: getStat(params.Stats.DEX),
			CON: getStat(params.Stats.CON),
			INT: getStat(params.Stats.INT),
			WIS: getStat(params.Stats.WIS),
			CHA: getStat(params.Stats.CHA),
			SIZ: getStat(params.Stats.SIZ),
		},
	}
	character.HP = getHP(character.Stats.CON, character.Stats.SIZ)
	character.DamageBonus = getDamageBonus(character.Stats.STR, character.Stats.SIZ)
	return character
}
