package helpers

func AbsInt(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
