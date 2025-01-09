package character

import (
	"math/rand"

	"github.com/ropfoo/gothulhu/internal/model"
)

func getRandomGender() model.Gender {
	if rand.Intn(2) == 0 {
		return "male"
	}
	return "female"
}
