package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("The server is listening on Port 3000")
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	listener2, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go acceptLoop(listener)
	acceptLoop(listener2)
}

func acceptLoop(l net.Listener) {
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New connection found!")
		go listenConnection(c)
	}

}
func listenConnection(conn net.Conn) {
	fmt.Println("Yay")
	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection has closed")
			return
		}

		//This is the message you received
		data := buffer[:dataSize]
		fmt.Print("Received message: ", string(data))
		f, err := os.OpenFile("database.txt", os.O_RDWR, 644)
		if err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		f.Write(data)
		// Send the message back
		_, err = conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print("Message sent: ", string(data))
	}
}
