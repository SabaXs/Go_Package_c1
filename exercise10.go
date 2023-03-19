package main

import "fmt"

var i int
var b bool
//a:=20
func main() {
	var s string
	fmt.Println("Default values of variables are: ")
	fmt.Printf("Default value of integers is: %v\nDefault value of boolean values is: %v\nDefault value of strings is: %v\n", i, b, s)
	a := 25
	fmt.Println("A variable declared by var can be accessed if it is outside the function as well, but a variable initialized with := cannot be accessed")
	fmt.Println(a)

}
