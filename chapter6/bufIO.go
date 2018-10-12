package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	file := os.Args[1]

	r, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {

		line := scanner.Text()
		if scanner.Err() != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}

		fmt.Println(line)
	}

}
