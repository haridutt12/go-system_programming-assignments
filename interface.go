package main

import (
	"fmt"
)


	type coordinates interface{
		Xaxix() int
		Yaxix() int
	}

	type points struct{
		X,Y int
	}

	// func (p points) Xaxix() int{
	// 	return p.X
	// }
	func  (p points) Xaxix() int{
		return p.X
	}

	func (p points) Yaxix() int{
		return p.Y
	}

	func findcoordinates(a coordinates){
		fmt.Printf("the coordinates are x=%d , y=%d",a.Xaxix(),a.Yaxix())
		fmt.Println()
	}

func main(){
	a := points{X:1, Y:2}
	fmt.Println(a)
	findcoordinates(a)

}
