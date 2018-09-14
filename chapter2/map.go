package main 
import (
	"fmt")
func main(){
	a := map[string]int{"sun":0,"mon":1,"tue":2}
	// a["sun"]=0
	fmt.Println(a)
	_,ok := a["tue"]
	if ok {
		fmt.Println("index exist")
	} else {
		fmt.Println("index tue doesnt exist")
	}
	fmt.Printf("size of map is %d",len(a))
	fmt.Println()
	for key,_ := range a{
		fmt.Println(key)
	}
}