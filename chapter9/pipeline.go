package main

import (
	"fmt"
	"os"
	"strconv"
)

func gennumbers(min, max int, out chan<- int) {
	fmt.Println("running channel 1")
	for i := min; i <= max; i++ {
		out <- i
	}
	close(out)
}

func findsquares(in <-chan int, out chan<- int) {
	fmt.Println("running channel 2")
	for i := range in {
		out <- i * i
	}
	close(out)
}

func squaresum(in <-chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	fmt.Println("square sum is :", sum)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	min, _ := strconv.Atoi(args[1])
	max, _ := strconv.Atoi(args[2])

	if min > max {
		fmt.Println("max should be greater then min")
		os.Exit(2)
	}
	c1 := make(chan int)
	c2 := make(chan int)

	go gennumbers(min, max, c1)
	go findsquares(c1, c2)
	squaresum(c2)

}
