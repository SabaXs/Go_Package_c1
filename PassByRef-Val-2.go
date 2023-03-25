package main

import "fmt"

func main() {
	a, b := "Sabahat", "Supriya"
	c, d := "Palak", "Bindu"

	fmt.Printf("Student a: %s, b: %s, c: %s, d: %s\n", a, b, c, d)
	fmt.Printf("(Memory) Address a: %p, b: %p, c: %p, d: %p\n", &a, &b, &c, &d)

	fmt.Println("\nUsing Swap method to differentiate between call by reference and call by value:")

	Swap(a, b)

	SwapRef(&c, &d)

	fmt.Printf("\nUpdated Values: \n")
	fmt.Printf("Student a:  %s, b: %s, c: %s, d: %s\n", a, b, c, d)
	fmt.Printf("(Memory) Address a: %p, b: %p, c: %p, d: %p\n", &a, &b, &c, &d)
}

func Swap(x, y string) {
	fmt.Printf("Swap by Value memory address: a:%p, b:%p\n", &x, &y)
	x, y = y, x
}

func SwapRef(x, y *string) {
	fmt.Printf("Swap by Reference memory address: c:%p, d:%p\n", x, y)
	*x, *y = *y, *x
}
