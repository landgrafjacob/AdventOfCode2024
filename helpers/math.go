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

func Mod(a, b int) int {
	return (a%b + b) % b
}

func AbsRune(a rune) rune {
	if a < 0 {
		return -1 * a
	} else {
		return a
	}
}
