package main 

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var s [3]string
	s[0] = "1 2 3 4 5"
	s[1] = "5 4 a b 1"
	s[2] = "6 7 2"

	argument := os.Args
	column, err := strconv.Atoi(argument[1])

	if err != nil {
		fmt.Println("error reading argument")
		os.Exit(-1)
	}
	if column < 0 {
		fmt.Println("enter valid column index")
		os.Exit(1)
	} 

		sum := 0
		for i := 0; i < len(s); i++ {
			data := strings.Fields(s[i])
		
				datanum, err := strconv.Atoi(data[column-1])
				if err != nil  {
					
					fmt.Println("not an integer",data[column-1])
				} else {
					sum = sum + datanum

				}
			
			
		}
			fmt.Println("summary : ", sum)
	}

	
