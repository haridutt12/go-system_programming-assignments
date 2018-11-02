package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func findIP(input string) string {

	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])"
	grammer := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchme := regexp.MustCompile(grammer)
	output := matchme.FindString(input)
	return output
}

func main() {

	args := os.Args

	if len(args) < 1 {
		fmt.Println("enter atleast 1 file as argument")
		os.Exit(1)
	}

	myIPs := make(map[string]int)
	for _, filename := range args[1:] {

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()

		r := bufio.NewReader(file)
		for {

			line, err := r.ReadString('\n')
			if err == io.EOF {
				fmt.Println(err)
				break
			} else if err != nil {
				fmt.Println(err)
				continue
			}
			ip := findIP(line)
			if myIPs[ip] >= 1 {
				myIPs[ip] = myIPs[ip] + 1
			} else {
				myIPs[ip] = 1
			}

		}

	}

	for key, val := range myIPs {

		fmt.Println(key, ":", val)
	}

}
