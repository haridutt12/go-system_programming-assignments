package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type File struct {
	Filename   string
	Lines      int
	Words      int
	Characters int
	Error      error
}

func processfiles(files []string, out chan<- File) {
	for _, filename := range files {
		var filetoprocess File
		filetoprocess.Characters = 0
		filetoprocess.Lines = 0
		filetoprocess.Filename = filename
		filetoprocess.Words = 0
		out <- filetoprocess
	}
	close(out)

}
func count(in <-chan File, out chan<- File) {

	for y := range in {
		filename := y.Filename
		f, err := os.Open(filename)
		if err != nil {
			y.Error = err
			out <- y
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
				y.Error = err
				out <- y
				continue
			}
			y.Lines = y.Lines + 1
			r := regexp.MustCompile("[^\\s]+")
			for range r.FindAllString(line, -1) {
				y.Words = y.Words + 1
			}
			y.Characters = y.Characters + len(line)
		}
		out <- y
	}
	close(out)
}
func calculate(in <-chan File) {
	var totalWords int = 0
	var totalLines int = 0
	var totalChars int = 0
	for x := range in {
		totalWords = totalWords + x.Words
		totalLines = totalLines + x.Lines
		totalChars = totalChars + x.Characters
		if x.Error == nil {
			fmt.Printf("\t%d\t", x.Lines)
			fmt.Printf("%d\t", x.Words)
			fmt.Printf("%d\t", x.Characters)
			fmt.Printf("%s\n", x.Filename)
		}
	}
	fmt.Printf("\t%d\t", totalLines)
	fmt.Printf("%d\t", totalWords)
	fmt.Printf("%d\ttotal\n", totalChars)
}
func main() {

	args := os.Args
	if len(args) == 1 {
		os.Exit(1)
	}
	files := make(chan File)
	values := make(chan File)
	go processfiles(args[1:], files)
	go count(files, values)
	calculate(values)
}
