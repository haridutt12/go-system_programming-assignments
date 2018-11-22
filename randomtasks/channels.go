package main

import (
	"fmt"
)

func sum(a []int, ch chan int) {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	ch <- s
}

func main() {
	a := []int{1, 2, -6, 5, 7, 8}
	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)

	A, b := <-ch, <-ch

	fmt.Println(A, b, A+b)

}
