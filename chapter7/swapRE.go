package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("expected 1 argument")
		os.Exit(1)
	}

	filename := args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	numoflines := 0
	numberOfLinesMatched := 0
	for {

		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file", err)
		}

		numoflines++
		r := regexp.MustCompile(`(.*)
		(\[\d\d\/(\w+)/\d\d\d\d:\d\d:\d\d:\d\d(.*)\]) (.*) (\d+)`)

		if r.MatchString(line) {
			numberOfLinesMatched++
			match := r.FindStringSubmatch(line)
			fmt.Println(match[1], match[6], match[5], match[2])

		}
	}

	fmt.Println("Line processed:", numoflines)
	fmt.Println("Line matched:", numberOfLinesMatched)

}
