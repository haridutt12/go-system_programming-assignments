package main

import (
	"fmt"
)

func writechannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)

}

func main() {

	c := make(chan int)
	x := 10
	go writechannel(c, x)

	// time.Sleep(2 * time.Second)
	fmt.Println("read:", <-c)

}
