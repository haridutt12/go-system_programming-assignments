package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil

}

func main() {

	args := os.Args
	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		filename := args[1]
		err := cat(filename)
		if err != nil {
			fmt.Println(err)
		}
	}

}
