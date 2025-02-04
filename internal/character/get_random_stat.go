package character

import "math/rand"

// get a random stat between 0 and 100
func getRandomStat() int {
	// Roll 3 times and average for a more bell-curve-like distribution
	rolls := 3
	sum := 0
	for i := 0; i < rolls; i++ {
		sum += rand.Intn(100)
	}
	return sum / rolls
}

func getStat(stat float32) []float32 {
	if stat > 0 {
		return []float32{float32(stat), float32(stat) / 2, float32(stat) / 5}
	}
	randomStat := getRandomStat()
	return []float32{float32(randomStat), float32(randomStat) / 2, float32(randomStat) / 5}
}
