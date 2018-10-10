package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	WalkFunction := func(path string, _ os.FileInfo, err error) error {
		_, err = os.Stat(path)
		if err != nil {
			return err
		}
		Basepath := filepath.Base(path)
		fmt.Println(Basepath)
		return nil

	}

	err = filepath.Walk(dir, WalkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
