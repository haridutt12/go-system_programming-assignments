package main

import (
	"fmt"
	"os"
)

func tripletToBinary(triplet string) string {

	if triplet == "rwx" {
		return "111"
	}
	if triplet == "-wx" {
		return "011"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "---" {
		return "000"
	}
	if triplet == "r-x" {
		return "101"
	}
	if triplet == "r--" {
		return "100"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "rw-" {
		return "110"
	}
	if triplet == "-w-" {
		return "010"
	}
	return "unknown"
}
func createbinary(binaryPermissions string) string {
	p1 := binaryPermissions[1:4]
	p2 := binaryPermissions[4:7]
	p3 := binaryPermissions[7:10]
	return tripletToBinary(p1) + tripletToBinary(p2) + tripletToBinary(p3)
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	filename := arguments[1]
	// file, _ := os.Open(filename)
	info, _ := os.Stat(filename)

	mode := info.Mode()

	fmt.Println("file mode is", mode)
	fmt.Println("as string is ", mode.String()[1:10])
	fmt.Println("binary is ", createbinary(mode.String()))

}
