package character

import "github.com/ropfoo/gothulhu/internal/model"

// generate a character with random stats
func GenerateCharacter(params model.CharacterParams) model.Character {
	var name string
	if params.Name == "" {
		name = GetName(params.Gender)
	} else {
		name = params.Name
	}

	var age int
	if params.Age == 0 {
		age = GetAge()
	} else {
		age = params.Age
	}

	character := model.Character{
		Name:   name,
		Age:    age,
		Gender: params.Gender,
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
