package helpers

func AbsInt(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func GCD(a, b int) int {
	tmpa := a
	tmpb := b
	for tmpb != 0 {
		tmpa, tmpb = tmpb, tmpa%tmpb
	}
	return tmpa
}
