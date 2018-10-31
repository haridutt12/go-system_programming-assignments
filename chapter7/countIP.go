package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	minusColumn := flag.Int("c", 1, "column number")
	flag.Parse()

	flags := flag.Args()

	if len(flags) < 1 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	column := *minusColumn

	MyIPs := make(map[string]int)

	for _, filename := range flags {

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()
		r := bufio.NewReader(file)
		// lines := r.ReadString("\n")
		// lines := strings.Split(string(r), "\n")
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				continue
			}
			data := strings.Fields(line)
			ip := data[column-1]
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}

			_, ok := MyIPs[ip]
			if ok {
				MyIPs[ip] = MyIPs[ip] + 1
			} else {
				MyIPs[ip] = 1
			}
		}

	}
	for key, val := range MyIPs {
		fmt.Println(key, val)
	}

}
