package main

import (
	"fmt"
	"sort"
)
type astruct struct {
	name string
	marks int
	cgpa float64
}

func main() {

	aslice := make([]astruct, 0)
	a := astruct{"ram", 80, 7.6}
	aslice = append(aslice, a)
	a = astruct{"shyam", 90, 8.6}
	aslice = append(aslice, a)
	a = astruct{"ghanshyam", 95, 8.7}
	aslice = append(aslice, a)

	fmt.Println("without sorting", aslice)

	sort.Slice(aslice, func(i, j int) bool {
		return aslice[i].cgpa < aslice[j].cgpa
		})

	fmt.Println("< :", aslice)

	bslice := []int{5, 6, 1, -3, 0}
	sort.Slice(bslice, func(i, j int) bool {
		return bslice[i] < bslice[j]
		})
	fmt.Println("ascending :", bslice)
}