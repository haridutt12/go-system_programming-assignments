package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunction(path string, info os.FileInfo, err error) error {

	_, err = os.Stat(path)

	if err != nil {
		return err
	}

	fmt.Println(path)
	return nil
}

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("enter valid no of arguments")
		os.Exit(1)
	}

	Path := args[1]

	err := filepath.Walk(Path, WalkFunction)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
