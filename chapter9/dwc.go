package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

func count(filename string) {

	var err error
	var numberOfLines int = 0
	var numberOfCharacters int = 0
	var numberOfWords int = 0
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
		}
		numberOfLines++
		r := regexp.MustCompile("[^\\s]+")
		for range r.FindAllString(line, -1) {
			numberOfWords++
		}
		numberOfCharacters += len(line)
	}
	fmt.Printf("\t%d\t", numberOfLines)
	fmt.Printf("%d\t", numberOfWords)
	fmt.Printf("%d\t", numberOfCharacters)
	fmt.Printf("%s\n", filename)
}

func main() {

	args := os.Args
	if len(args) == 1 {
		os.Exit(1)
	}
	var waitgroup sync.WaitGroup
	for _, filename := range args[1:] {
		waitgroup.Add(1)
		go func(filename string) {
			count(filename)
			defer waitgroup.Done()
		}(filename)
	}
	waitgroup.Wait()
}
