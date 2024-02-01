package piscine

func ListSize(l *List) int {
	i := 0
	take := l.Head
	for take != nil {
		take = take.Next
		i++
	}
	return i
}
