package doublelinkedlist

type Elem struct {
	next, prev *Elem

	list *List

	Value interface{}
}

func (e *Elem) Next() *Elem {
	if p := e.next; e.list != nil && p != &e.list.keyElem {
		return p
	}
	return nil
}

func (e *Elem) Prev() *Elem {
	if p := e.prev; e.list != nil && p != &e.list.keyElem {
		return p
	}
	return nil
}

type List struct {
	keyElem Elem
	len     int
}

func (l *List) Init() *List {
	l.keyElem.next = &l.keyElem
	l.keyElem.prev = &l.keyElem
	l.len = 0
	return l
}

func New() *List { return new(List).Init() }

func (l *List) Len() int { return l.len }

func (l *List) Front() *Elem {
	if l.len == 0 {
		return nil
	}
	return l.keyElem.next
}

func (l *List) Back() *Elem {
	if l.len == 0 {
		return nil
	}
	return l.keyElem.prev
}

func (l *List) zeroList() {
	if l.keyElem.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Elem) *Elem {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Elem) *Elem {
	return l.insert(&Elem{Value: v}, at)
}

func (l *List) remove(e *Elem) *Elem {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) move(e, at *Elem) *Elem {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

func (l *List) Remove(e *Elem) interface{} {
	if e.list == l {

		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v interface{}) *Elem {
	l.zeroList()
	return l.insertValue(v, &l.keyElem)
}

func (l *List) PushBack(v interface{}) *Elem {
	l.zeroList()
	return l.insertValue(v, l.keyElem.prev)
}

func (l *List) InsertBefore(v interface{}, mark *Elem) *Elem {
	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Elem) *Elem {
	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}

func (l *List) MoveToFront(e *Elem) {
	if e.list != l || l.keyElem.next == e {
		return
	}

	l.move(e, &l.keyElem)
}

func (l *List) MoveToBack(e *Elem) {
	if e.list != l || l.keyElem.prev == e {
		return
	}

	l.move(e, l.keyElem.prev)
}

func (l *List) MoveBefore(e, mark *Elem) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

func (l *List) MoveAfter(e, mark *Elem) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l *List) PushBackList(other *List) {
	l.zeroList()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.keyElem.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.zeroList()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.keyElem)
	}
}
