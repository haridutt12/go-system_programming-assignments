package main

import (
	"fmt"
	"time"
)

func routine1() {

	time.Sleep(2 * time.Second)
	fmt.Println("riunning routine 1")
}

func main() {

	fmt.Println("starting main routine")
	var a int
	b, _ := fmt.Scanln(&a)
	go routine1()
	fmt.Println(b)
	go func() {

		fmt.Println("running routine 2")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("exiting")
}
