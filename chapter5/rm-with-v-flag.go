package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	details := flag.Bool("-v", false, "--verbose")
	flag.Parse()
	args := os.Args

	if len(args) == 1 {

		fmt.Println("please enter argument")
		os.Exit(1)
	}

	file := args[1]
	err := os.Remove(file)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
	if *details {

		fmt.Println("removed :", args)
	}

}
