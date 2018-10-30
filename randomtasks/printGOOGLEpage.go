package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	home := os.Args[1]
	resp, err := http.Get("https://www.google.com/")

	if err != nil {

		fmt.Println("error fatching url", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buf := make([]byte, 100000)

	ioutil.WriteFile(home, body, 0644)
	destData, err := os.Open(home)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, _ = destData.Read(buf)
	// a := string(body)
	fmt.Printf("%s", buf)

}
