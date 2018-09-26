package main 

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Data int
	Left *Tree
	Right *Tree
}

func create(n int) *Tree{
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i:=0; i<2*n; i++ {
		ran := rand.Intn(n)
		t = insert(t, ran)
	}
	return t
}
func insert(node *Tree, d int) *Tree {

	if node == nil {

		return  &Tree{d, nil, nil}
	}

	if d == node.Data {

		return node
	}

	if d < node.Data {

		node.Left = insert(node.Left, d)
		return node
	}

	node.Right =  insert(node.Right, d)
	return node


}

func traverse(node *Tree) {

	if node != nil {

		fmt.Print(" ",node.Data)
	    traverse(node.Left)
		traverse(node.Right)
	}
}

func main() {

	tree := create(30)
	traverse(tree)
	fmt.Println("value of root is :", tree.Data)

}
