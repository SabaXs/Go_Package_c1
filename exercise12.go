package main

import "fmt"

func ctf(temp int) int {
	return (temp * 9 / 5) + 32
}

func ftc(far float32) float32 {
	return (far - 32) * 5 / 9
}

func main() {
	fmt.Println("The Temperature 23 degrees Celsius in Fahrenheit is : ")
	fmt.Println(ctf(23))
	fmt.Println("The Temperature 90 degrees Fahrenheit in Celsius is : ")
	fmt.Println(ftc(90))
}
