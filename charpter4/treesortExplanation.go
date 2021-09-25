// As we all know, binary sort tree has three traversal method.
// Inorder,preorder,postorder
package main

import "fmt"

type tree struct {
	v           int
	left, right *tree
}

func main() {
	fmt.Println("Let's read the code.")
}

func inOrder(t *tree) {
	if t.left != nil {
		inOrder(t.left)
	}
	fmt.Println(t.v)
	if t.right != nil {
		inOrder(t.right)
	}
}

func preOrder(t *tree) {
	fmt.Println(t.v)
	if t.left != nil {
		preOrder(t.left)
	}
	if t.right != nil {
		preOrder(t.right)
	}
}

func postOrder(t *tree) {
	if t.left != nil {
		postOrder(t.left)
	}
	if t.right != nil {
		postOrder(t.right)
	}
	fmt.Println(t.v)
}
