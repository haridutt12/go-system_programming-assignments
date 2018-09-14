package main 

import (
	"fmt"
	// "math/max"
	)

func printslice(x []int){

	for _,num := range x {
		fmt.Printf("%d",num)
	}
	fmt.Println()
}

func main(){

	var inp int
	fmt.Scanln(&inp)
	aslice := []int{inp}
    fmt.Println("enter 0 to terminate")

	for  inp != 0 {
		fmt.Scanln(&inp)
		aslice = append(aslice,inp)
	} 

    max := 0
    min := 100000

	for  i:=0; i< len(aslice); i++ {
		if aslice[i] > max {
			max = aslice[i]
		}
		if aslice[i] < min {
			min = aslice[i]

        }
	}

	fmt.Printf("max: %d and min: %d", max,min)
	fmt.Println()

}