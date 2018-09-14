package main 

import (
	"fmt"
)


func printslice(x []int){
	for _,num := range x {
		fmt.Printf("%d",num)
	}
	fmt.Println()
}

func main(){
	aslice := []int{1,2,3,4}
	bslice := make([]int, 5, 10)
    copy(bslice,aslice)
	// println("count:", n)
	fmt.Println(bslice)


}
