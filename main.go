package main

import (
	"fmt"
)

func main(){

	var x int = 15

	var y float32 = float32(x)
	fmt.Printf("%v, %T\n",x,x)
	fmt.Printf("%v, %T\n",y,y)
}