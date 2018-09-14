package main 

import (
	"fmt"
    "os"
    "strconv"
	)

func main(){
    
    fmt.Println("enter 0 to terminate")
    arguments := os.Args
    a, _ := strconv.Atoi(arguments[1])
	aslice := []int{ a }
	for i:=2; i<len(arguments); i++ {
			a, _ := strconv.Atoi(arguments[i]) 
			if a != 0 {
            aslice = append(aslice, a)
            } else {
            	break
            }

    }

     max := aslice[0]
     min := aslice[0]

	for  i:=1; i< len(aslice); i++ {
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