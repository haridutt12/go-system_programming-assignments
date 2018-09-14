package main
import (
	"fmt"
)
func main(){
	var a int 
	println("enter an integer value")
	fmt.Scanln(&a)
	b := a*10
	fmt.Println("successfully scanned :",b)
}
