package main 
import (
	"fmt")

func withpointer(x *int){
	*x = *x * *x
}

type complex struct{
	x, y int
}

func newcomplex(x, y int ) *complex{
	return &complex{x, y}
}

func main(){
 x := -4
 withpointer(&x)
 fmt.Println(x)
 w := newcomplex(4, -5)
 fmt.Println(*w)
 fmt.Println(w)

}