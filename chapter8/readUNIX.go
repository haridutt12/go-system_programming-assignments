package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func readunix(r io.Reader) {

	buf := make([]byte, 1024)
	for {

		n, _ := r.Read(buf[:])
		fmt.Println("read", string(buf[0:n]))
	}
}

func main() {

	c, _ := net.Dial("unix", "/tmp/aSocket.sock")
	defer c.Close()
	go readunix(c)
	n := 0
	for {
		message := []byte("hi reading " + strconv.Itoa(n) + "\n")
		_, _ = c.Write(message)
		time.Sleep(5)
		n += 1
	}

}
