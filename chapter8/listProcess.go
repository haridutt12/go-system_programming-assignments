package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	PS, err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(PS)

	args := []string{"ps", "-a", "-x"}

	env := os.Environ()

	err = syscall.Exec(PS, args, env)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
