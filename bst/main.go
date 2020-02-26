package main

import "fmt"

type node struct {
	left  *node
	right *node
	val   int
}

func addNode(curr *node, child *node) (root node) {
	if curr != nil {
		if child.val < curr.val {
			if curr.left != nil {
				addNode(curr.left, child)
			} else {
				curr.left = child
			}
		} else {
			if curr.right != nil {
				addNode(curr.right, child)
			} else {
				curr.right = child
			}
		}
	}

	return *curr
}

func createNode(value int) (curr node) {
	return node{
		left:  nil,
		right: nil,
		val:   value,
	}
}

func printInOrder(curr *node) {

	if curr.left != nil {
		printInOrder(curr.left)
	}

	fmt.Print("|")
	fmt.Print(curr.val)

	if curr.right != nil {
		printInOrder(curr.right)
	}

}

func main() {
	children := []node{createNode(1), createNode(6), createNode(2)}
	root := createNode(5)
	root.left = &children[0]
	root.right = &children[1]
	root.left.right = &children[2]

	printInOrder(&root)
}
