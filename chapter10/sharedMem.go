package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var write = make(chan int)
var read = make(chan int)

func writechannel(val int) {
	write <- val
}
func readchannel() int {
	return <-read
}

func monitor() {
	var value int
	for {
		select {
		case newvalue := <-write:
			value = newvalue
			fmt.Println(value)
		case read <- value:
		}
	}
}

func main() {

	rand.Seed(time.Now().Unix())
	go monitor()
	var waitgroup sync.WaitGroup
	for i := 0; i < 20; i++ {
		waitgroup.Add(1)
		go func() {
			defer waitgroup.Done()
			writechannel(rand.Intn(100))
		}()
	}
	waitgroup.Wait()
	fmt.Printf("\nLast value: %d\n", readchannel())
}
