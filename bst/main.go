package main

import (
	"fmt"
)

type node struct {
	left *node
	right *node
	val int
}

func addNode(curr *node, child *node) (root *node){
	if curr != nil {
		if child.val < curr.val {
			if curr.left != nil {
				addNode(curr.left, child)
			} else {
				curr.left = child;
			}
		} else {
			if curr.right != nil {
				addNode(curr.right, child)
			} else {
				curr.right = child;
			}
		}
	}

	return curr
}

func createNode(value int) (curr node){
	return node{
		left: nil,
		right: nil,
		val: value,
	}
}

func printInOrder(root *node){
	curr := root
	stack := []node{}

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, *curr)
			curr = curr.left
		}

		curr, stack = &stack[0], stack[1:]
		fmt.Print("|")
		fmt.Print(curr.val)
		curr = curr.right
	}

	fmt.Println("|");
	return
}

func main() {
	root := createNode(5);
	children := []node{};
	children = append(children,createNode(1));
	children = append(children,createNode(5));
	children = append(children,createNode(7));
	children = append(children,createNode(8));
	children = append(children,createNode(3));

	for i := 0; i < len(children); i++ {
		root = *addNode(&root, &children[i])
	}

	printInOrder(&root);
}