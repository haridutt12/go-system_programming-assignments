package main 

import (
	"log"
)

func main(){
	x := 1
	log.Println(x)
	x+=1
	log.Println(x)
	x+=1
	log.Panicf("%d",x)
}