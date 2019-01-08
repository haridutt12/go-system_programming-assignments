package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatalln(err)
		}

		for {
			buffer := make([]byte, 1400)
			dataSize, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("The connection has closed!")
				return
			}

			data := buffer[:dataSize]
			fmt.Println("Received message: ", string(data))
			break
		}

	}
}
