package main

import (
	"fmt"
	"os"
	"strconv"
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
	p1 := binaryPermissions[0:3]
	p2 := binaryPermissions[3:6]
	p3 := binaryPermissions[6:9]
	p1 = tripletToBinary(p1)
	p2 = tripletToBinary(p2)
	p3 = tripletToBinary(p3)
	p1Int, _ := strconv.ParseInt(p1, 2, 64)
	p2Int, _ := strconv.ParseInt(p2, 2, 64)
	p3Int, _ := strconv.ParseInt(p3, 2, 64)
	returnValue := p1Int*100 + p2Int*10 + p3Int
	tempReturnValue := int(returnValue)
	returnString := "0" + strconv.Itoa(tempReturnValue)
	return returnString
}

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	filename := args[1]
	permission := args[2]

	if len(permission) != 9 {
		fmt.Println("permission should be 9 characters example : rwxrwxrwx")
		os.Exit(1)
	}

	BinaryPermission := createbinary(permission)
	newPerm, _ := strconv.ParseUint(BinaryPermission, 0, 32)
	newMode := os.FileMode(newPerm)
	os.Chmod(filename, newMode)

}
