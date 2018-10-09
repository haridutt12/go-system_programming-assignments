package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	Test := flag.Bool("t", false, "testing only")
	flag.Parse()

	args := flag.Args()
	source := args[0]
	destination := args[1]

	if len(args) == 0 || len(args) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	permission := os.ModePerm

	_, err := os.Stat(destination)

	if os.IsNotExist(err) {
		os.MkdirAll(destination, permission)
	} else {
		fmt.Println(destination, " already exist")
		os.Exit(1)
	}

	WalkFunction := func(currentpath string, _ os.FileInfo, err error) error {

		fileinfo, _ := os.Lstat(currentpath)
		if fileinfo.Mode()&os.ModeSymlink != 0 {
			fmt.Println(currentpath, "skipping")
			return nil
		}
		fileinfo, err = os.Stat(currentpath)

		if err != nil {
			fmt.Println("*", err)
			return err
		}
		mode := fileinfo.Mode()
		if mode.IsDir() {
			tempPath := strings.Replace(currentpath, source, "", 1)
			pathToCreate := destination + "/" + filepath.Base(source) +
				tempPath
			if *Test {
				fmt.Println(":", pathToCreate)
				return nil
			}
			_, err := os.Stat(pathToCreate)
			if os.IsNotExist(err) {
				os.MkdirAll(pathToCreate, permission)
			} else {
				fmt.Println("Did not create", pathToCreate, ":", err)
			}
		}
		return nil
	}
	err = filepath.Walk(source, WalkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
