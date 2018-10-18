package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func count(filename string) (int, int, int) {

	countline := 0
	countchar := 0
	countwords := 0
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer data.Close()
	r := bufio.NewReader(data)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Print("error reading file", err)
			break
		}
		countline++
		rxp := regexp.MustCompile("[^\\s]+")
		for range rxp.FindAllString(line, -1) {
			countwords++
		}
		countchar += len(line)
	}
	return countline, countwords, countchar
}

func main() {

	args := flag.Args()

	minusl := flag.Bool("l", false, "count_lines")
	minusw := flag.Bool("w", false, "count_words")
	minusc := flag.Bool("c", false, "count_characters")
	flag.Parse()

	if len(args) == 1 {
		fmt.Println("enter minimum 1 file")
		os.Exit(1)
	}
	printAll := false
	totalLines := 0
	totalChar := 0
	totalWords := 0

	for _, file := range flag.Args() {

		line, words, chars := count(file)
		totalLines += line
		totalWords += words
		totalChar += chars

		if (*minusc && *minusl && *minusw) || (!*minusc && !*minusl && !*minusw) {

			fmt.Println(line, words, chars, file)
			printAll = true
			continue
		}

		if *minusl {
			fmt.Print(line)
		}

		if *minusw {
			fmt.Print(words)
		}

		if *minusc {
			fmt.Print(chars)
		}

		fmt.Println(file)
	}
	if printAll && (len(args) != 1) {
		fmt.Println(totalLines, totalWords, totalChar, "total")
	}
	if len(args) != 1 && *minusl {
		fmt.Print(totalLines)
	}
	if len(args) != 1 && *minusw {
		fmt.Print(totalWords)
	}
	if len(args) != 1 && *minusc {
		fmt.Print(totalChar)
	}

}
