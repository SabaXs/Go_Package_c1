package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
  return x-y
}

func mul(x int, y int) int {
  return x*y
}

func main() {
	fmt.Println(add(42, 13))
  fmt.Println(sub(12,4))
  fmt.Println(mul(2,6))
}
