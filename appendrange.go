package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var mid []int
	for i := min; i < max; i++ {
		mid = append(mid, i)
	}
	return mid
}
