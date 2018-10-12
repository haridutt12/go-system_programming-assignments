package main

import (
	"fmt"
	"io"
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

	buf := make([]byte, 8)
	_, err = io.ReadFull(r, buf)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(os.Stdout, string(buf))
	fmt.Println()

}
