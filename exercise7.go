package main

import (
	"fmt"
	"math/cmplx"
)

var (
	s  string     = "data types example"
	i  int8       = 126
	i2 int16      = 1000
	bl bool       = true
	f  float32    = 19.47
	f2 float32    = 3.4e+38
	m  uint64     = 1<<64 - 1
	z  complex128 = cmplx.Sqrt(-8 + 10i)
)

func main() {
  fmt.Printf("The different data types of golang: ")
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", i2, i2)
	fmt.Printf("Type: %T Value: %v\n", bl, bl)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", f2, f2)
	fmt.Printf("Type: %T Value: %v\n", m, m)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
