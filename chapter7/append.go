package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("invalid no of arguments enter filename and message")
		os.Exit(1)
	}

	filename := os.Args[1]
	message := os.Args[2]

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(file, "%s\n", message)

}
