package main

import "fmt"

func passbyVal(a int) {
	a++
	fmt.Println("Inside Pass By Value Function: a =", a)
}
func passbyRef(b *int) {
	(*b)++
	fmt.Println("Inside Pass By Reference Function: b =", *b)
}

func main() {
	a := 5
	fmt.Println("Before Pass By Value: a =", a) 
	passbyVal(a)
  fmt.Println("After Pass By Value: a =", a)

	b := 5
	fmt.Println("Before Pass By Reference: b =", b)
	passbyRef(&b)
	fmt.Println("After Pass By Reference: b =", b)
}
