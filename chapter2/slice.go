package main 
import (
	"fmt")
func change(x []int){
	x[2]=-2
}

func printslice(x []int){
	for _,num := range x {
		fmt.Printf("%d",num)
	}
	fmt.Println()
}

func main(){
	aslice := []int{1,2,3,4,5}
	fmt.Println("slice before chaanges")
	printslice(aslice)
	fmt.Println("slice after chaanges")
	change(aslice)
	printslice(aslice)
	aslice = append(aslice, -100)
	// anotherslice := []int{6,5,4,3,2,1}
	// aslice = copy(aslice,anotherslice)
	printslice(aslice)
	fmt.Printf("after capacity: %d, length: %d\n",cap(aslice),len(aslice))
}