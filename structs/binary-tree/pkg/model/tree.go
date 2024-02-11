package model

import "fmt"

type Node struct {
	Right *Node
	Left  *Node

	Val int
}

// This insert has binary search tree logic.
func (n *Node) Insert(val int) {
	if n == nil {
		return
	} else if val <= n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func PrintNode(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.Val)

	PrintNode(n.Left)
	PrintNode(n.Right)
}

func (n *Node) NumberOfNodes() int {
	if n == nil {
		return 0
	}

	return n.Left.NumberOfNodes() + n.Right.NumberOfNodes() + 1
}

func (n *Node) Depth() int {
	depth := 0

	if n == nil {
		return depth
	}

	if n.Left.Depth() > n.Right.Depth() {
		depth = n.Left.Depth()
	} else {
		depth = n.Right.Depth()
	}

	return depth + 1
}

func PrintTree(n *Node) {
	if n != nil {
		PrintTree(n.Right)
		PrintTree(n.Left)

		fmt.Printf("%d ", n.Val)
	}
}

func InvertTree(n *Node) {
	if n != nil {
		n.Right, n.Left = n.Left, n.Right

		InvertTree(n.Left)
		InvertTree(n.Right)
	}
}
