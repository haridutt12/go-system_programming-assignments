package main

import (
	"fmt"
	"time"
)

func gennumbers(out chan<- int) {
	for i := 0; i < 10; i++ {
		select {
		case out <- i:
		default:
			fmt.Println("channel capacity full")
			break
		}
	}
}

func main() {

	ch := make(chan int, 5)
	gennumbers(ch)
	for i := 0; i < 10; i++ {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
			fmt.Println("channel capacity  h full")
			break
		}
	}

	time.Sleep(2 * time.Second)

}
