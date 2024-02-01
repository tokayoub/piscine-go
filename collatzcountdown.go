package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	i := 0
	for start != 1 {
		i++
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
	}
	return i
}
