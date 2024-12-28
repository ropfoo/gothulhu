package character

type DamageBonus struct {
	Damage  string `json:"damage"`
	Stature string `json:"stature"`
}

var damageBonus = map[int]DamageBonus{
	-2: {Damage: "-2", Stature: "-2"},
	-1: {Damage: "-1", Stature: "-1"},
	0:  {Damage: "0", Stature: "0"},
	1:  {Damage: "1d4", Stature: "+1"},
	2:  {Damage: "1d6", Stature: "+2"},
	3:  {Damage: "2d6", Stature: "+3"},
	4:  {Damage: "3d6", Stature: "+4"},
	5:  {Damage: "4d6", Stature: "+5"},
	6:  {Damage: "5d6", Stature: "+6"},
}

func getDamageBonus(str int, siz int) DamageBonus {
	sum := str + siz
	switch {
	case sum <= 64:
		return damageBonus[-2]
	case sum <= 84:
		return damageBonus[-1]
	case sum <= 124:
		return damageBonus[0]
	case sum <= 164:
		return damageBonus[1]
	case sum <= 204:
		return damageBonus[2]
	case sum <= 284:
		return damageBonus[3]
	case sum <= 364:
		return damageBonus[4]
	case sum <= 444:
		return damageBonus[5]
	default:
		return damageBonus[6]
	}
}
