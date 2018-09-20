package main

import (
	"fmt"
)

type Node struct {

	Value int
	Next *Node
	
}

func addnode(n *Node, m int) int {

	if root == nil {
		n = &Node{m, nil}
		root = n
	    return 0
	}

	if m == n.Value {
		fmt.Println("node already exist :", m)
		return -1
	}

	if n.Next == nil {
		n.Next = &Node{m, nil}
		return -2
	}
	return addnode(n.Next, m)
}

func traverse(n *Node) {
	if n ==  nil {
		fmt.Println("empty list")
	}
	for n != nil {
		fmt.Printf( "%d -->", n.Value)
		n = n.Next
	}
	fmt.Println()
}

func addbetween(n *Node, m int) int {

	if n.Value == 3 {

		temp := n.Next
		t := &Node{m, nil}
		n.Next = t
		t.Next = temp
		return 0

	}
	return addbetween(n.Next, m)
}

func deletenode(n *Node, m int) int {

	p := n.Next

	if n.Value == m && n == root {
		root = n.Next
		return 0
	}

	if p.Value == m && p.Next == nil {
		n.Next = nil
		return 1
	}

	if p.Value == m {
		n.Next = p.Next
		return 2
	}

	return deletenode(n.Next, m)
}

var root = new(Node)

func main() {
 	
 	fmt.Println(root)
 	root = nil
 	traverse(root)
	addnode(root, 1)
	addnode(root, 1)
	traverse(root)
	addnode(root, 2)
	addnode(root, 3)
	addnode(root, 4)
	addbetween(root, 6)
	traverse(root)
	deletenode(root, 6)
	traverse(root)

}