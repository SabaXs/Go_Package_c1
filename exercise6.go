package main

import "fmt"

var age, grade int = 19, 9

func main() {
	var name, course, year = "Sabahat Shaik", "CSE", "Third"
	fmt.Printf("The following are the details of the student: \nName: %v \nAge: %v \nYear: %v\nCourse: %v\nGrade: %v", name, age, year, course, grade)
}
