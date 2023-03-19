package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Sabahat", "Shaik")
	fmt.Println(a, b)
	c := "RHS"
	d := "LHS"
	e, f := swap(c, d)
	fmt.Printf("If %v = %v", c, d)
	fmt.Printf(" then %v = %v", e, f)
}
