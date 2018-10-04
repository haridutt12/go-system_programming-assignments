package main 
import (
	"fmt"
	"os"
)
func main() {

	args := os.Args

	if len(args) == 1 {

		fmt.Println("please enter argument")
		os.Exit(1)
	}

	file := args[1]
	err := os.Remove(file)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

}