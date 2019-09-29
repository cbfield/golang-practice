package main

import (
	"fmt"
	"strconv"
)

func main(){

	var x int = 15
	var y string = strconv.Itoa(x)

	fmt.Printf("%v, %T\n",x,x)
	fmt.Printf("%v, %T\n",y,y)
}