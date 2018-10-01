package main 

import (

	"fmt"
)
func quicksort(a []int, low, high int ) {

	pivot := a[high]

	for i:=0; i< len(a); i++ {

		if a[i] >= pivot {
			a[i], pivot = pivot, a[i]
			quicksort(a, low, i)
			quicksort(a, i, high)
			break
		}

	}
}

func main() {
 		
	arr := []int{2, -3, 8, 5, -7, 10}
	quicksort(arr, 0, len(arr) - 1 )
	fmt.Println(arr)

}