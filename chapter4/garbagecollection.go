package main 

import (

	"fmt"
	"runtime"
	"time"
)

func printstats(mem runtime.MemStats) {

	runtime.ReadMemStats(&mem)
	fmt.Println("Memory Alloc :", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {

	var mem runtime.MemStats
	printstats(mem)

	for i:=0; i<10; i++ {
		a := make([]byte, 100000000)
		if a == nil {
			fmt.Println("Operation failed!")
			}
	}
	printstats(mem)

	for i:=0; i<10; i++ {
		a := make([]byte, 100000000)
		if a == nil {
			fmt.Println("Operation failed!")
			}
		time.Sleep(5*time.Second)
		// printstats(mem)
	}
	printstats(mem)


}