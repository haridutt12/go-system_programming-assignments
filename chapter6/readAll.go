package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	src := args[1]
	dest := args[2]
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(dest, data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
