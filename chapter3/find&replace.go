package main 

import (
	"fmt"
	"regexp"
)

func main() {

	var s [3]string
	s[0] = "1 2 3 4 5"
	s[1] = "5 4 a b 1"
	s[2] = "6 7 2"

	parse, _ := regexp.Compile("1")
 	
 	for i:=0; i<len(s); i++ {
 	 	res :=	parse.ReplaceAllString(s[i], "harry")
 	 	fmt.Println(res)
	}


}