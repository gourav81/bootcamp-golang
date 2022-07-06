// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Printf(node.data)
}

func main() {

	root := &Node{"+", nil, nil}
	root.left = &Node{"a", nil, nil}
	root.right = &Node{"-", nil, nil}
	root.right.left = &Node{"*", nil, nil}
	root.right.right = &Node{"c", nil, nil}
	root.right.left.left = &Node{"*", nil, nil}
	root.right.left.right = &Node{"^", nil, nil}
	root.right.left.left.left = &Node{"d", nil, nil}
	root.right.left.left.right = &Node{"e", nil, nil}
	root.right.left.right.left = &Node{"f", nil, nil}
	root.right.left.right.right = &Node{"g", nil, nil}
	fmt.Println("PreOrder Traversal")
	preOrder(root)
	fmt.Println("\nPostOrder Traversal")
	postOrder(root)

}
