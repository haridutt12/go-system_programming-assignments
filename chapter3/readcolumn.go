package main 

import (
	"fmt"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "Harry is a nice guy"
	s[1] = "loves to eat"
	s[2] = "dema is a pretty girl"

	colm := 1

	for i:=0 ; i<len(s) ;i++ {
		
         data := strings.Fields(s[i])
         if len(data) >= colm {
         	fmt.Println(data[colm-1])
         }

		
	}
}