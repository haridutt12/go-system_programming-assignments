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

	pid := os.Getpid()
	fmt.Println(pid)
	ch := make(chan os.Signal, 1)

	signal.Notify(ch)
	go func() {

		chval := <-ch
		switch chval {
		case syscall.SIGTERM:
			handlesignal(chval)
			os.Exit(-1)
		case os.Interrupt:
			handlesignal(chval)
		default:
			fmt.Println("ignoring : ", chval)
		}

	}()

	for {

		fmt.Println("dema")
		time.Sleep(2 * time.Second)
	}

}
