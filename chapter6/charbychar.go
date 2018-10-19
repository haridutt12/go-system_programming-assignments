package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	file := args[1]
	r, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(1)
	}
	in := string(r)
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		fmt.Print(s.Text())
	}

}
