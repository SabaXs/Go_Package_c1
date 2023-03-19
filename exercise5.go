package main

import "fmt"

/*func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}*/

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

func split3(sum int) (x, y, z int) {
	x = sum / 3
	y = sum - 2*x
	z = sum - x - y
	return
}

func main() {
	fmt.Println("The split of number 23 is: ")
	fmt.Println(split(23))
	fmt.Println("Splitting 1234 in 2 halves gives: ")
	fmt.Println(split(1234))
	fmt.Println("Splitting the number 123 in 3 halves: ")
	fmt.Println(split3(123))
}
