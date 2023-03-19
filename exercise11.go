package main

import "fmt"

func main() {
	var a, b, c int = 18, 17, 19
	var sub1, sub2, sub3 string = "Math", "Science", "English"
	fmt.Println("The following are the grades of a student: ")
	fmt.Printf("%v : %v\n%v : %v\n%v : %v\n", sub1, a, sub2, b, sub3, c)
}
