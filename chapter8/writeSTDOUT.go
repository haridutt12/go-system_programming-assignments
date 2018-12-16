package main

import (
	"io"
	"os"
)

func main() {

	args := os.Args
	mystring := ""
	if len(args) == 1 {
		mystring = "Invalid No of Arguments"
	} else {

		mystring = args[1]
	}

	io.WriteString(os.Stdout, mystring)
	io.WriteString(os.Stdout, "\n")
}
