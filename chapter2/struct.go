package main 

import (
	"fmt"
	// "reflect"
)

func main(){

	type message struct{
		X, Y int
		Label string
	}
	p1 := message{1,2,"a message"}
	p2 := message{}
	p2.Label = "message2"
	fmt.Println(p1)
	fmt.Println(p2)
}