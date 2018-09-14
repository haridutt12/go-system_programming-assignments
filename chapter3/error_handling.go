package main 

import (
	"fmt"
	"errors"
	"log"
)
func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("cannot be divided by is zero")
    }
    if x%y !=0 {
    	remainder := errors.New("there is a remainder")
    	return x/y, remainder, nil
    } else {
    	return x/y, nil, nil
    }
}

func main(){
	var a, b int
	fmt.Println("enter divident and divisor")
	fmt.Scanln(&a, &b)
	result, rem, err := division(a,b)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
	if rem != nil {
		fmt.Println(rem)
	}
}