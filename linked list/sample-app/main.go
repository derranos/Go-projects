package main

import (
	"fmt"
	"testMod/sample-app/listfuncs"
)

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
	l := new(listfuncs.List)
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
			l.AddElement(n)
			fmt.Printf("number added (%d)\n", n)
		case '-':
			if l.Head == nil {
				fmt.Print("list is empty\n")
			} else {
				l.DelElement()
				fmt.Print("last number deleted\n")
			}
		case 'l':
			fmt.Printf("the length of the list is %d\n", l.LenOfList())
		case 'd':
			var pos int
			fmt.Scan(&pos)
			if l.DelAtPos(pos) {
				fmt.Print("the number was successfully deleted\n")
			} else {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		case 'c':
			var pos, num int
			fmt.Scan(&pos, &num)
			if l.ChangeAtPos(pos, num) {
				fmt.Print("the number was successfully changed\n")
			} else {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		case 'g':
			var pos int
			fmt.Scan(&pos)
			if !l.GetAtPos(pos) {
				fmt.Print("the length of the list is less than the specified position or pos < 0\n")
			}
		default:
			fmt.Print("error key\n")
		}
	}
}
