package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	minusL := flag.Int("l", 1, "length")
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	filename := os.Args[1]

	linenumber := *minusL

	file, err := ioutil.ReadFile(filename)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)

	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		lines[i] = fmt.Sprintf("%d: %s ", linenumber, line)
		linenumber -= 1
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(filename, []byte(output), 0664)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("pricessed", linenumber-*minusL, "lines!")

}
