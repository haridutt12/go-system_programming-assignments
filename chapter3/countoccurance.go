package main 

import (
	"fmt"
	"strings"
)
func main() {

	var s [3]string
	s[0] = "sale sale sale"
	s[1] = "bumper sale 50 percent off"
	s[2] = "on kids and men clothing"

	amap := make(map[string]int)

	for i:=0; i<len(s); i++ {
		data := strings.Fields(s[i])
		for j := 0; j < len(data); j++ {
			amap[data[j]] += 1
		}
	}

	for key, val := range amap {
		fmt.Println(key,"=>",val)
	}

}