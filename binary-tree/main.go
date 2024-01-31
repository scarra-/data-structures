package main

import (
	"fmt"

	"github.com/aadejanovs/binary-tree/pkg/model"
)

func main() {
	//      1
	//     / \
	//    2   3
	//  /  \   \
	// 4    5   6

	rootNode := &model.Node{Val: 1}

	rootNode.Left = &model.Node{Val: 2}
	rootNode.Right = &model.Node{Val: 3}
	rootNode.Right.Right = &model.Node{Val: 6}
	rootNode.Left.Left = &model.Node{Val: 4}
	rootNode.Left.Right = &model.Node{Val: 5}

	// rootNode.Insert(2)
	// rootNode.Insert(3)
	// rootNode.Insert(4)
	// rootNode.Insert(5)
	// rootNode.Insert(6)

	// model.PrintNode(rootNode)

	fmt.Println(rootNode.NumberOfNodes())
	fmt.Println(rootNode.Depth())

	model.PrintTree(rootNode)

	model.InvertTree(rootNode)
	fmt.Println()
	model.PrintTree(rootNode)

	// a.Left = b
	// a.Right = c
	// b.Left = d
	// b.Right = e
	// c.Right = f

	//      a
	//     / \
	//    b   c
	//  /  \   \
	// d    e   f
}

// func depthFirstValues(root *model.Node) {
// 	stack := []string{}
// 	for {
// 		if len(stack) <= 0 {
// 			return
// 		}

// 	}
// }
