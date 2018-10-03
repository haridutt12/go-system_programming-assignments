package main

import (
	"fmt"
	"os"
	"strings"
	"flag"

)

func main() {

	Minusa := flag.Bool("a", false, "a")
	flag.Parse()

	flags := flag.Args()

	if len(flags) == 0 {

		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	found := false
	file := flags[0]
	path := os.Getenv("PATH")

	pathSlice := strings.Split(path, ":")

	for _, directory := range pathSlice {

		fullpath := directory + "/" + file
		fileinfo, err := os.Stat(fullpath)

		if err == nil {

			mode := fileinfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					found = true
					if *Minusa == true {
						fmt.Println(fullpath)
						} else {
						fmt.Println(fullpath)
						os.Exit(0)
						}
				}
			}
		}

	}
	if found == false {
		os.Exit(1)
	}


}
