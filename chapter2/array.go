package main  
import(
"fmt")
func main(){
	
	// var a int 
	// fmt.Scanln(&a)
	var myarray [2]int
	fmt.Println("enter the elements of array")
	for i:=0; i<len(myarray); i++{
		fmt.Scanln(&myarray[i])
	}
	fmt.Println(myarray[0])
}