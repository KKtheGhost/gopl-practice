// This go program will be written in kivinsae's standard style
// This program shows a way to sort numbers by a binary tree with recursion structure.
package main

import "fmt"

type binaryTree struct {
	intValue          int
	leftPos, rightPos *binaryTree
}

func main() {
	a := []int{14, 42, 622, 254, 443, 55}
	sortInteger(a)
	fmt.Println(a)
}

func sortInteger(intValues []int) {
	var binaryTreeRoot *binaryTree
	for _, v := range intValues {
		binaryTreeRoot = addInteger(binaryTreeRoot, v)
	}
	appendInteger(intValues[:0], binaryTreeRoot)
}

// Recursion append function
func appendInteger(intValues []int, pTree *binaryTree) []int {
	if pTree != nil {
		intValues = appendInteger(intValues, pTree.leftPos)
		intValues = append(intValues, pTree.intValue)
		intValues = appendInteger(intValues, pTree.rightPos)
	}
	return intValues
}

// Recursion add integer function
func addInteger(pTree *binaryTree, intValue int) *binaryTree {
	if pTree == nil {
		pTree = new(binaryTree)
		pTree.intValue = intValue
		return pTree
	}
	if intValue < pTree.intValue {
		pTree.leftPos = addInteger(pTree.leftPos, intValue)
	} else {
		pTree.rightPos = addInteger(pTree.rightPos, intValue)
	}
	return pTree
}
