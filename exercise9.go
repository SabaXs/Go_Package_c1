package main

import "fmt"

const (
	age = 10
)

func after(x int) int   { return x + 5 }
func brother(x int) int { return x * 2 }
func b_after(x int) int { return (x * 2) + 5 }

func main() {
	fmt.Println("Seema's age now is: ", age)
	fmt.Println("Her brother Raju is twice as old as her. He is now: ", brother(age))
	fmt.Println("After 5 years Seema would be: ", after(age))
	fmt.Println("At the same time, her brother Raju would be: ", b_after(age))
}
