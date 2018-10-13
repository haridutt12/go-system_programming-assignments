package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	file := os.Args[1]

	dest, err := os.Create(file)
	if err != nil {
		fmt.Println("os.Create : ", err)
		os.Exit(1)
	}

	defer dest.Close()

	fmt.Fprint(dest, "i m testing fmt.Fprinff() to write this text into a file,so that it is in formatted form")
}
