package improc

func Clamp(val int) int {
	if val < 0 {
		return 0
	}
	if val > 255 {
		return 255
	}
	return val
}
