package model

import "encoding/json"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Character struct {
	Name        string      `json:"name"`
	Age         int         `json:"age"`
	Gender      Gender      `json:"gender"`
	HP          int         `json:"hp"`
	DamageBonus DamageBonus `json:"damage_bonus"`
	Stats       Stats       `json:"stats"`
	Skills      []Skill     `json:"skills"`
}

type DamageBonus struct {
	Damage string `json:"damage"`
	Build  string `json:"build"`
}

type Stats struct {
	STR []float32 `json:"str"`
	DEX []float32 `json:"dex"`
	CON []float32 `json:"con"`
	INT []float32 `json:"int"`
	WIS []float32 `json:"wis"`
	CHA []float32 `json:"cha"`
	SIZ []float32 `json:"siz"`
}

type Skill struct {
	Name  string    `json:"name"`
	Score []float32 `json:"score"`
}

func (c Character) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(json)
}
