package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var Bufsize int

func Copy(src, dest string, Bufsize int) error {
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
	buff := make([]byte, Bufsize)
	destination, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer destination.Close()
	for {
		data, err := file.Read(buff)
		if err != nil && err != io.EOF {
			return err
		}
		if data == 0 {
			break
		}
		if _, err := destination.Write(buff[:data]); err != nil {
			return err
		}

	}
	return err
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	src := args[1]
	dest := args[2]
	Bufsize, _ = strconv.Atoi(args[3])
	err := Copy(src, dest, Bufsize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("copied")
	}

}
