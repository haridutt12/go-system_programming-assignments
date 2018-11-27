package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handlesignal(sig os.Signal) {
	fmt.Println("got", sig)
}

func main() {

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	go func() {

		for {

			chval := <-ch
			fmt.Println(chval)
			handlesignal(chval)
		}
	}()

	for {

		fmt.Println("harry")
		time.Sleep(2 * time.Second)
	}
}
