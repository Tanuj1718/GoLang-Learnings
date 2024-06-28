package main

import (
	"fmt"
	my_package "mylearning/my_package"
)

func main() {
	fmt.Println("Learn GO")
	my_package.PrintMessage("File structure testing")
	var name string = "Hello"
	fmt.Println(name)
}
