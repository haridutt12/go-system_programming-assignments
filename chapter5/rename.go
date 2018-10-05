package main 

import (

	"fmt"
	"flag"
	"os"
	"path/filepath"
)

func main() {

	minusO := flag.Bool("o", false, "o")
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {

		fmt.Println("enter 2 arguments")
		os.Exit(1)
	}

	src := args[0]
	dest := args[1]

	srcinfo, err := os.Stat(src)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	} else {

		mode := srcinfo.Mode()
		if mode.IsRegular() == false {

			fmt.Println("Sorry! we support regular files only")
		}
	}

	newdest := dest
	destinfo, err := os.Stat(dest)

	if err == nil {
		mode := destinfo.Mode()
		if mode.IsDir() {
			justsrcname := filepath.Base(src)
			newdest = dest + "/" + justsrcname
		}
		dest = newdest
	}

	destinfo, err = os.Stat(dest)

	if err == nil {
		if *minusO == false {
			fmt.Println("file alraedy exist")
			os.Exit(1)
		}
	}else {
		err = os.Rename(src, dest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}