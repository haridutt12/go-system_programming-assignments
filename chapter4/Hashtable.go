package main 

import "fmt"

type Node struct {

	Val int
	Next *Node
}

type Hash struct {

	Table map[int]*Node
	Size int
}

func hashfunction(i, s int) int {

	return i%s
}

func insert(hash *Hash, val int) int {

	index := hashfunction(val, hash.Size)
	data := Node{val, Hash.Table[index]}
	Hash.Table[index] = &data
	return index
}