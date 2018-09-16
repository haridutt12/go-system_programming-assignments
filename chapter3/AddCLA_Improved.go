package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	args := os.Args
	sum := 0
	for i := 1; i < len(args); i++ {
		temp, err := strconv.Atoi(args[i])
		if err == nil{
			sum += temp
			
		} else {
			fmt.Println("ignoring argument :",args[i])
		}
    }
    fmt.Printf("sum of arguments is : %d", sum)
    fmt.Println()
}