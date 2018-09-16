package main 

import (
	"fmt"
	"regexp"
)

func main() {
	s1 := "Hari dutt parashar"
	s2 := "Hari"
	res, _ := regexp.MatchString(s2, s1)
	fmt.Println(res)
    
    parse, err := regexp.Compile("[Hh]ari")
    if err != nil {
    	fmt.Println(err)
    } else {
    	fmt.Println(parse.MatchString("hari"))
    	fmt.Println(parse.MatchString("Hari"))
    	fmt.Println(parse.MatchString("H ari"))
    	fmt.Println(parse.ReplaceAllString("hari dutt", "HARI"))
    }
}
