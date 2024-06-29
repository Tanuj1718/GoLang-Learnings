package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Printf("data type of numbers : %T\n", numbers)
	fmt.Println("Length: ", len(numbers))
	numbers = append(numbers,3,44,44)
	fmt.Println("Length: ", len(numbers))

}