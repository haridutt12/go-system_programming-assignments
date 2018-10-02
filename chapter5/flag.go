package main 

import (
	"fmt"
	"flag"
)

func main() {

	minusO := flag.Bool("o", false, "o")
	minusI:= flag.Bool("i", false, "i")
	minusK := flag.Int("k", 0, "k")

	flag.Parse()

 	fmt.Println("-o:", *minusO)
	fmt.Println("-i:", *minusI)
	fmt.Println("-K:", *minusK)

	for index, key := range flag.Args() {
		fmt.Println("index:", index, "key :", key)
	}

}