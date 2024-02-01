package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	str := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = str
		l.Tail = str
		return
	}
	l.Tail.Next = str
	l.Tail = str
}
