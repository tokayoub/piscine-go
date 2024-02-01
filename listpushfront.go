package piscine

func ListPushFront(l *List, data interface{}) {
	str := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = str
	} else {
		str.Next = l.Head
	}
	l.Head = str
}
