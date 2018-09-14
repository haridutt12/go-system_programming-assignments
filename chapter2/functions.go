package main 
import (
	"fmt"
	"math")
func squreroot(a float64) float64{
	return math.Sqrt(a)
}
func main(){
	var c float64
	fmt.Scanln(&c)
	b := squreroot(c)
	fmt.Println("square root of",c ,"is ",b)
}