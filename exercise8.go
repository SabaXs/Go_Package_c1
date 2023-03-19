package main

import "fmt"

const age = 19

func main() {
	const name = "Sabahat"
	fmt.Printf("Hello, my name is %v. I am %v years old\n", name, age)
	ans := true
	fmt.Println("Ans is not a const function: ", ans)
	ans = false
	fmt.Println("A constant function can be changed: ", ans)
}
