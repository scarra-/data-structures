package main

import (
	"fmt"
	"time"
)

type Node struct {
	Val int

	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n.Val > val {
		if n.Left != nil {
			n.Left.Insert(val)
		} else {
			n.Left = &Node{Val: val}
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(val)
		} else {
			n.Right = &Node{Val: val}
		}
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if n.Val > val {
		return n.Left.Search(val)
	} else if n.Val < val {
		return n.Right.Search(val)
	}

	return true
}

func main() {
	tree := &Node{Val: 100}

	tree.Insert(200)
	tree.Insert(50)
	tree.Insert(70)
	tree.Insert(80)
	tree.Insert(24)
	tree.Insert(18)
	tree.Insert(233)
	tree.Insert(245)
	tree.Insert(80)
	tree.Insert(201)
	tree.Insert(1)
	tree.Insert(873)
	tree.Insert(9)
	tree.Insert(1000)

	start := time.Now()

	fmt.Println(tree.Search(245))

	fmt.Println("Time elapsed: ", time.Since(start))
}
