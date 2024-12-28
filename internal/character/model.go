package character

import "encoding/json"

type Character struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	STR  int    `json:"str"`
	DEX  int    `json:"dex"`
	CON  int    `json:"con"`
	INT  int    `json:"int"`
	WIS  int    `json:"wis"`
	CHA  int    `json:"cha"`
}

func (c Character) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(json)
}
