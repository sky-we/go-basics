package fmt

func TooLarge(n int) bool {
	if n >= 1000 {
		return true
	}
	return false
}
