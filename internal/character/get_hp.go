package character

func getHP(con float32, siz float32) int {
	return int((con + siz) / 10)
}
