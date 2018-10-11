package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}

	num, _ := strconv.ParseInt(os.Args[1], 10, 64)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, num)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d is %x in Little Endian", num, buf)
	buf.Reset()

	err = binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d is %x in Big Endian", num, buf)

}
