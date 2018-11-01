package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	flags := os.Args
	// if len(flags) < 2 {
	// 	fmt.Println("enter atleast 1 file as argument")
	// 	os.Exit(1)
	// }

	for _, filename := range flags[1:] {

		info, _ := os.Stat(filename)
		fmt.Printf("%+v\n", info.Sys())
		fmt.Println(info.Sys().(*syscall.Stat_t).Uid)
		fmt.Println(info.Sys().(*syscall.Stat_t).Gid)

	}

}
