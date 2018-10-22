package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func random(min, max int) int {

	return rand.Intn(max-min) + min

}

var mu sync.Mutex

func writedatatofile(i int, file *os.File, w *sync.WaitGroup) {
	mu.Lock()
	time.Sleep(time.Duration(random(10, 1000)) * time.Millisecond)
	fmt.Fprintf(file, "From %d, writing %d\n", i, 2*i)
	fmt.Printf("writing %d to file\n", i)
	w.Done()
	mu.Unlock()
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("enter 1 argument")
		os.Exit(1)
	}
	number := 3

	filename := os.Args[1]
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	var w *sync.WaitGroup = new(sync.WaitGroup)
	w.Add(number)

	for i := 0; i < number; i++ {
		writedatatofile(i, file, w)
	}

	w.Wait()

}
