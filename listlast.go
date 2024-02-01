package piscine

func ListLast(l *List) interface{} {
	var a interface{} = nil
	take := l.Head
	for take != nil {
		if take.Next == nil {
			a = take.Data
		}
		take = take.Next
	}
	return a
}
