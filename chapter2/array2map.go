package main 

import (
	"fmt"
	"strconv"
)

func arraytomap(arr [4]int) map[string]int{
	    amap := make(map[string]int)
		for i := 0; i< len(arr); i++{
		amap[strconv.Itoa(arr[i])] = arr[i]
	}
	return amap
}
func main(){
	arr := [4]int{1,2,3,4}
	bmap := arraytomap(arr)
	
	for key,val := range bmap{
		fmt.Println(key,val)
	}

}

