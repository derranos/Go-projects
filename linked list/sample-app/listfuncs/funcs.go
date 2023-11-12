package listfuncs

import "fmt"

type Node struct {
	num  int
	next *Node
}
type List struct {
	Head *Node
}

func (l *List) AddElement(i int) {
	newL := new(Node)
	newL.num = i
	if l.Head == nil {
		l.Head = newL
		return
	}
	cur := l.Head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newL
}
func (l *List) DelElement() {
	if l.Head.next == nil {
		l.Head = nil
		return
	}
	cur := l.Head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
}
func (l *List) LenOfList() int {
	cur := l.Head
	if cur == nil {
		return 0
	}
	k := 1
	for cur.next != nil {
		cur = cur.next
		k++
	}
	return k
}

func (l *List) DelAtPos(pos int) bool {
	if pos < 0 || l.Head == nil {
		return false
	}
	if pos == 0 {
		l.Head = l.Head.next
		return true
	}
	k := 0
	cur := l.Head
	for cur.next != nil && k < pos-1 {
		k++
		cur = cur.next
	}
	if cur.next == nil {
		return false
	}
	cur.next = cur.next.next
	return true
}
func (l *List) ChangeAtPos(pos, num int) bool {
	if pos < 0 {
		return false
	}
	k := 0
	cur := l.Head
	for k < pos && cur != nil {
		k++
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	cur.num = num
	return true
}
func (l *List) GetAtPos(pos int) bool {
	if pos < 0 {
		return false
	}
	k := 0
	cur := l.Head
	for k < pos && cur != nil {
		cur = cur.next
		k++
	}
	if cur == nil {
		return false
	}
	fmt.Printf("%d\n", cur.num)
	return true
}
