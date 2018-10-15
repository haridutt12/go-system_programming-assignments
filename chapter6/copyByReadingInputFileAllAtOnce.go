package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dest string) (int64, error) {
	fileinfo, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !fileinfo.Mode().IsRegular() {
		fmt.Println("not a regular file")
		os.Exit(1)
	}
	file, _ := os.Open(src)
	defer file.Close()
	destination, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer destination.Close()
	bytes, err := io.Copy(destination, file)
	return bytes, err
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	src := args[1]
	dest := args[2]
	bytes, err := Copy(src, dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("copied", bytes, "bytes")
	}

}
