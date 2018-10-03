package main 
import (
	"fmt"
	"os"
	"path/filepath"

)

func main() {

	args := os.Args
	dir, err := os.Getwd()

	if err == nil {

		fmt.Println(dir)
	} else {

		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) == 1 {

		return 
	}

	if args[1] != "-P" {

		return 
	}

	fileinfo, err := os.Lstat(dir)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	} else {

		if fileinfo.Mode() & os.ModeSymlink !=0 {

			realpath, _ := filepath.EvalSymlinks(dir)
			fmt.Println("Path : ", realpath)
		}
		
	}

}