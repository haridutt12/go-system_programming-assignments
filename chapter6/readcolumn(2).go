package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	colm := flag.Int("col", 1, "column")
	flag.Parse()

	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("invalid no of arguments add atleast one filename")
		os.Exit(1)
	}

	column := *colm

	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	for _, filename := range flags {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("cannot open file", filename)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}

		}

	}

}
