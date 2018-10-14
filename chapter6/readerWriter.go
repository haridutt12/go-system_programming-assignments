package main

import (
	"fmt"
	"io"
	"os"
)

func countchar(w io.Reader) int {
	total := 0
	buff := make([]byte, 16)
	for {
		n, err := w.Read(buff)
		if err != nil && err != io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}

		total += n
	}
	return total
}

func writenumofchars(w io.Writer, x int) {
	fmt.Fprintf(w, "%d\n", x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		os.Exit(-1)
	}
	defer f.Close()
	count := countchar(f)
	filename = filename + ".Count"
	f, err = os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	defer f.Close()
	writenumofchars(f, count)

}
