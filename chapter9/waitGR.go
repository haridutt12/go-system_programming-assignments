package main

import (
	"fmt"
	"sync"
)

func main() {

	// var i int64
	var waitgroup sync.WaitGroup
	waitgroup.Add(10)
	for i := 0; i < 10; i++ {

		go func(x int) {
			defer waitgroup.Done()
			fmt.Println("running routine %d", x)
		}(i)
	}
	waitgroup.Wait()
	fmt.Println("exiting main routine")

}
