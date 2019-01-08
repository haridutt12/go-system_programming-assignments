package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func excludeNames(name string, exclude string) bool {
	if exclude == "" {
		return false
	}
	if filepath.Base(name) == exclude {
		return true
	}
	return false
}

func excludeExtensions(name string, extension string) bool {
	if extension == "" {
		return false
	}
	basename := filepath.Base(name)
	s := strings.Split(basename, ".")
	length := len(s)
	basenameExtension := s[length-1]
	if basenameExtension == extension {
		return true
	}
	return false
}

func main() {

	excludeSocket := flag.Bool("s", false, "Sockets")
	excludePipes := flag.Bool("p", false, "Pipes")
	excludeSymLinks := flag.Bool("sl", false, "Symbolic Links")
	excludeDirectories := flag.Bool("d", false, "Directories")
	excludeFiles := flag.Bool("f", false, "Files")
	excludeSpecificFile := flag.String("x", "", "Files")
	excludeExtention := flag.String("ext", "", "Extensions")

	flag.Parse()
	flags := flag.Args()

	printAll := false
	if *excludeSocket && *excludePipes && *excludeSymLinks && *excludeDirectories && *excludeFiles {
		printAll = true
	}

	if !(*excludeSocket || *excludePipes || *excludeSymLinks || *excludeDirectories || *excludeFiles) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := flags[0]

	walkFunction := func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if excludeNames(path, *excludeSpecificFile) {
			return nil
		}

		if excludeExtensions(path, *excludeExtention) {
			return nil
		}

		if printAll == true {
			fmt.Println(path)
			return nil
		}

		mode := fileInfo.Mode()
		if mode.IsRegular() && *excludeFiles {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *excludeDirectories {
			fmt.Println(path)
			return nil
		}

		// os.Lstat() function that gives you information about a file or directory and
		// the use of the Mode() function on the return value of the os.Lstat() call in order to
		// compare the outcome with the os.ModeSymlink constant, which is the symbolic link bit
		fileInfo, _ = os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *excludeSymLinks {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *excludePipes {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *excludeSocket {
				fmt.Println(path)
				return nil
			}
		}

		return nil
	}

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
