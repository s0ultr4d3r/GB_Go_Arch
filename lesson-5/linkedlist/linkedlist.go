package linkedlist

type Node struct {
	next *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Len() *Node {
	return l.len
}

func (l *List) Find(elem int) *Node {
	if l.Head != nil {
		for tmp := l.Head; tmp.Next != nil; tmp = tmp.Next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Add(prev *Node, node *Node) {
	l.Len++
	if l.Head == nil {
		l.Head = node
		return
	}
	node.Next = prev.Next
	prev.Next = node
}

func (l *List) Append(node *Node) {
	l.Add(l.Tail, node)
}

func (l *List) Delete(node *Node) {

}

//14113
