package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("provide valid no of arguments")
		os.Exit(1)
	}
	file := args[1]
	aByteslice := []byte("Haridutt Parashar")
	ioutil.WriteFile(file, aByteslice, 0644)

	f, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	anotherbyteslice := make([]byte, 100)
	n, err := f.Read(anotherbyteslice)

	fmt.Printf("Read %d Bytes %s ", n, anotherbyteslice)

}
