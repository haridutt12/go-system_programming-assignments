package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func read(s string) []byte {

	var r io.Reader
	r, err := os.Open(s)
	check(err)
	p := make([]byte, 8)
	_, _ = r.Read(p)

	return p
}

func write(s string, i []byte) {

	file, err := os.OpenFile(s, os.O_RDWR, 0644)

	check(err)
	defer file.Close()

	file.WriteAt(i, 0) // Write at 0 beginning

}

func main() {

	args := os.Args
	cloud, err := strconv.Atoi(args[1])
	check(err)
	switch cloud {
	case 1:
		key := read("aws.txt")
		data := int32(binary.BigEndian.Uint32(key))
		fmt.Println("key for aws :", data)
		write("aws.txt", key)

	case 2:

		// key := read("gcp.txt")
		// data := binary.BigEndian.Uint16(key)
		// fmt.Println("key for gcp :", data+1)
		// write("gcp.txt", key)
		// write("aws.txt", key)
		// io.write

	case 3:

		// key := read("alicloud.txt")
		// data := binary.BigEndian.Uint16(key)
		// fmt.Println("key for alicloud :", data+1)
		// write("alicloud.txt", key)

	}

}
