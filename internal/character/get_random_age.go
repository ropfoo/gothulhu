package character

import (
	"math/rand"
	"time"
)

func getRandomAge() int {
	// Create a local random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate base age between 20-40 as most likely range
	baseAge := r.Intn(21) + 20 // 20-40

	// 30% chance to add 10-35 years for older characters
	if r.Float32() < 0.3 {
		baseAge += r.Intn(26) + 10
	}

	return baseAge
}
