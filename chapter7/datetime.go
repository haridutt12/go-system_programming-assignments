package main

import (
	"fmt"
	"time"
)

func main() {

	epochtime := time.Now().Unix()
	fmt.Println("unix time", epochtime)

	t := time.Now()
	t.Format(time.RFC3339)

	fmt.Println(t, t.Format(time.RFC3339))

	fmt.Println(t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)

	t1 := time.Now()

	fmt.Println("time difference ", t1.Sub(t))

	loc, _ := time.LoadLocation("Europe/London")

	londontime := t.In(loc)

	fmt.Println("london: ", londontime)

}
