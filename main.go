package main

import (
	"fmt"
)

func main(){

	var x int = 8 //1000
	var y int = 2 //0010

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	fmt.Println()

	fmt.Println(x & y) 	//0000
	fmt.Println(x | y) 	//1010
	fmt.Println(x ^ y)	//1010
	fmt.Println(x &^ y)	//1000

	fmt.Println()

	fmt.Println(x << 3)
	fmt.Println(x >> 3)
}