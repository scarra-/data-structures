package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l *LinkedList) Append(node *Node) {
	l.head.Append(node)
	l.length++
}

func (l *LinkedList) DeleteSingleOccurrenceByValue(val int) {
	if l.head == nil {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	traverser := l.head

	for traverser.next != nil {
		if traverser.next.data == val {
			traverser.next = traverser.next.next
			l.length--
			continue
		}

		traverser = traverser.next
	}
}

func (l *LinkedList) DeleteAllOccurrencesByValue(val int) {
	for l.head != nil && l.head.data == val {
		l.head = l.head.next
		l.length--
	}

	traverser := l.head
	var prev *Node

	for traverser != nil {
		if traverser.data == val {
			if prev != nil {
				prev.next = traverser.next
				l.length--
			}

			traverser = traverser.next
		} else {
			prev = traverser
			traverser = traverser.next
		}
	}
}

func (n *Node) PrintData() {
	if n != nil {
		fmt.Printf("%d ", n.data)
		n.next.PrintData()
	}
	fmt.Println()
}

func (n *Node) Append(newNode *Node) {
	if n.next == nil {
		n.next = newNode
		return
	}

	n.next.Append(newNode)
}

func (l *LinkedList) Print() {
	if l.head != nil {
		l.head.PrintData()
	}
}

func (l LinkedList) PrintWithLoop() {
	n := l.head

	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", n.data)
		n = n.next
	}

	fmt.Println()
}

func (l LinkedList) Values() []int {
	result := make([]int, l.length)

	n := l.head

	for i := 0; i < l.length; i++ {
		result[i] = n.data
		n = n.next
	}

	return result
}

func (l *LinkedList) HasLoop() bool {
	if l.head == nil {
		return false
	}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	llist := LinkedList{}
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 1})
	llist.Prepend(&Node{data: 2})
	llist.Prepend(&Node{data: 3})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 5})

	llist.PrintWithLoop()

	// llist.DeleteSingleOccurrenceByValue(5)
	// fmt.Println(llist.Values())

	llist.Append(&Node{data: 9})

	// llist.DeleteAllOccurrencesByValue(4)
	llist.PrintWithLoop()
}
