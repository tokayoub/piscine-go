package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		nextNodeL := current.Next
		current.Next = nil
		current = nextNodeL
	}
	l.Head = nil
}
