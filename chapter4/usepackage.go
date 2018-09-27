package main

import (
	"fmt"
	"mypackage"
)
func init() {
	fmt.Println("this function is called by default at first")
}

func main() {

	temp := mypackage.Add(5, 10)
    fmt.Println(temp)
    fmt.Println(mypackage.Pi)
}