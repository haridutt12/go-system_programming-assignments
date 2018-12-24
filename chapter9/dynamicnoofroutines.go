package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	// var i int64
	args := os.Args
	Nroutines, _ := strconv.Atoi(args[1])

	var waitgroup sync.WaitGroup
	// waitgroup.Add(Nroutines)
	for i := 0; i < Nroutines; i++ {
		waitgroup.Add(1)
		go func(x int) {
			defer waitgroup.Done()
			fmt.Println("running routine ", x)
		}(i)
	}
	waitgroup.Wait()
	fmt.Println("exiting main routine")

}
