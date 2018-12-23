package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("starting main routine")

	for i := 0; i < 10; i++ {

		go func(x int) {
			time.Sleep(1)
			fmt.Println("running routine", x)
		}(i)

	}

	time.Sleep(1 * time.Second)

	fmt.Println("exiting")
}
