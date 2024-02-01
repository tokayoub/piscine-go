package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	mid := make([]int, max-min)
	for i, n := 0, min; n < max; i, n = i+1, n+1 {
		mid[i] = n
	}
	return mid
}
