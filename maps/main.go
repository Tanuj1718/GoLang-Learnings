package main

import "fmt"

func main() {
	// key <-> value
	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["a"] = 45
	studentGrades["b"] = 56
	fmt.Println(studentGrades["a"])
}