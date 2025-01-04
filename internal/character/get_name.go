package character

import (
	"math/rand"

	"github.com/ropfoo/gothulhu/internal/data"
	"github.com/ropfoo/gothulhu/internal/model"
)

func GetName(gender model.Gender) string {
	surname := data.Surnames[rand.Intn(len(data.Surnames))]

	var firstName string
	if gender == model.Male {
		firstName = data.MaleNames[rand.Intn(len(data.MaleNames))]
	} else {
		firstName = data.FemaleNames[rand.Intn(len(data.FemaleNames))]
	}
	return firstName + " " + surname
}
