package main

import (
	"flag"
	"fmt"
	"os"
)

func addvalues(buf *[]byte, count int) {
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {

		intByte := byte(i)
		*buf = append(*buf, intByte)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("provide 1 argument")
		os.Exit(1)
	}

	capacity := flag.Int("cap", 0, "buffer capacity")
	number := flag.Int("n", 0, "iterations")
	flag.Parse()

	if *capacity <= 0 || *number <= 0 {
		fmt.Println("provide valid no of flags with positive values")
		os.Exit(1)
	}

	filename := os.Args[1]

	_, err := os.Stat(filename)

	if err == nil {
		fmt.Println("file already exists")
		os.Exit(1)
	}

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buf := make([]byte, *capacity)

	for i := 0; i < *capacity; i++ {
		addvalues(&buf, *number)
		_, err := file.Write(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		buf = nil
	}

}
