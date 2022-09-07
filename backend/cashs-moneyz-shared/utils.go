package cashs_moneyz_shared

func IAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func IMax(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func IMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
