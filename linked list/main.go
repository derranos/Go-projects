// linked list project main.go
package main

import (
	"fmt"
)

type Node struct {
	num  int
	next *Node
}
type List struct {
	head *Node
}

func (l *List) func0(i int) {
	newL := new(Node)
	newL.num = i
	if l.head == nil {
		l.head = newL
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newL
	}
}
func (l *List) func1() {
	if l.head.next == nil {
		l.head = nil
	} else {
		cur := l.head
		for cur.next.next != nil {
			cur = cur.next
		}
		cur.next = nil
	}
}
func (l *List) func2() int {
	cur := l.head
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

func (l *List) func3(pos int) bool {
	if pos < 0 || l.head == nil {
		return false
	}
	if pos == 0 {
		l.head = l.head.next
		return true
	}
	k := 0
	cur := l.head
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
func (l *List) func4(pos, num int) bool {
	if pos < 0 {
		return false
	}
	k := 0
	cur := l.head
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
func (l *List) func5(pos int) bool {
	if pos < 0 {
		return false
	}
	k := 0
	cur := l.head
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
func main() {
	//LINKED-LIST
	fmt.Print(
		`to add an element, enter + and a number
to remove an item from the end of the list, type -
to find out the size of the list, enter l
to delete an item at the "i" position, type d and "i"
to change the element to the "i" position, enter c and "i" and the number
to get an element at position "i", enter g and "i"
to finish the job, enter 0`)
	fmt.Print("\n")
	l := new(List)
	for true {
		fmt.Print("enter the command\n")
		var s string
		fmt.Scan(&s)
		c := rune(s[0])
		switch c {
		case '0':
			fmt.Print("bye!")
			return
		case '+':
			var n int
			fmt.Scan(&n)
			l.func0(n)
			fmt.Printf("number added (%d)\n", n)
		case '-':
			if l.head == nil {
				fmt.Print("list is empty\n")
			} else {
				l.func1()
				fmt.Print("last number deleted\n")
			}
		case 'l':
			fmt.Printf("the length of the list is %d\n", l.func2())
		case 'd':
			var pos int
			fmt.Scan(&pos)
			if l.func3(pos) {
				fmt.Print("the number was successfully deleted\n")
			} else {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		case 'c':
			var pos, num int
			fmt.Scan(&pos, &num)
			if l.func4(pos, num) {
				fmt.Print("the number was successfully changed\n")
			} else {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		case 'g':
			var pos int
			fmt.Scan(&pos)
			if !l.func5(pos) {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		default:
			fmt.Print("error key\n")
		}
	}
}
