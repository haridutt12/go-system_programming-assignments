package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
)

var BUFFER int64
var FILESIZE int64
var BYTESWRITTEN int64 = 0

func IntermediateCopyingStats() {

	progress := float64(BYTESWRITTEN) / float64(FILESIZE) * 100
	fmt.Printf("Progress: %.2f%%\n", progress)

}

func Copy(src, dest string, buf, size int64) error {

	source, err := os.Open(src)
	if err != nil {
		return err
	}

	defer source.Close()

	buffer := make([]byte, buf)

	_, err = os.Stat(dest)

	if err == nil {
		return fmt.Errorf("File %s already exists.", dest)
	}
	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	for {

		n, err := source.Read(buffer)

		if err != nil && err != io.EOF {

			return err
		}

		if n == 0 {
			break
		}
		if _, err := destination.Write(buffer[:n]); err != nil {
			return err
		}
		BYTESWRITTEN = BYTESWRITTEN + int64(n)

	}
	return err

}

func main() {

	args := os.Args
	if len(args) != 4 {
		fmt.Println("invalid no of arguments")
		os.Exit(1)
	}
	src := args[1]
	srcStat, err := os.Stat(src)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !srcStat.Mode().IsRegular() {
		fmt.Println("not a regular file source ")
		os.Exit(1)
	}

	dest := args[2]

	BUFFER, _ = strconv.ParseInt(args[3], 10, 64)

	FILESIZE = srcStat.Size()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {

		for {

			sig := <-sigs
			switch sig {

			case os.Interrupt:
				IntermediateCopyingStats()

			default:
				fmt.Println("ignored :", sig)

			}

		}
	}()

	fmt.Printf("copying %s to %s\n", src, dest)

	err = Copy(src, dest, BUFFER, FILESIZE)

	if err != nil {
		fmt.Println("error copying file")
	}
}
