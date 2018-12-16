package main

import (
	"bufio"
	"fmt"
	"os"
)

var f *os.File

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		f = os.Stdin
	} else {
		filenamme := arguments[1]
		file, err := os.Open(filenamme)
		if err != nil {
			fmt.Println(err)
		}
		f = file
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
