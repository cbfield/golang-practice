package main

import (
	"fmt"
)

var (
	actor string = "Elijah Wood"
	location string = "New York"
	isGood bool = true
)
func main(){

	// var num1 int
	// num1 = 1

	// var num2 int = 2
	
	// num3 := 3

	fmt.Printf("%v, %T\n", actor,actor)
	fmt.Printf("%v, %T\n", location,location)
	fmt.Printf("%v, %T\n", isGood,isGood)
}