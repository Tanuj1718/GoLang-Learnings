package main

import "fmt"

func main() {
	var num int
	num = 2

	var ptr *int //int is written because we are storing address of interger in ptr
	ptr = &num
	fmt.Println(ptr, "->", *ptr)

	num1 := 5
	ptr1 := &num1
	fmt.Println(ptr1, "->", *ptr1)
}