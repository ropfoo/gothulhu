package model

import "encoding/json"

type Character struct {
	Name        string      `json:"name"`
	Age         int         `json:"age"`
	HP          int         `json:"hp"`
	DamageBonus DamageBonus `json:"damage_bonus"`
	Stats       Stats       `json:"stats"`
}

type DamageBonus struct {
	Damage  string `json:"damage"`
	Stature string `json:"stature"`
}

type Stats struct {
	STR int `json:"str"`
	DEX int `json:"dex"`
	CON int `json:"con"`
	INT int `json:"int"`
	WIS int `json:"wis"`
	CHA int `json:"cha"`
	SIZ int `json:"siz"`
}

func (c Character) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(json)
}
