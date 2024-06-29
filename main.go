package main

import (
	"bufio"
	"fmt"
	my_package "mylearning/my_package"
	"os"
)

func main() {
	fmt.Println("Learn GO")
	my_package.PrintMessage("File structure testing")
	var name string = "Hello"
	fmt.Println(name)
	var age int;
	fmt.Scan(&age);
	fmt.Println(age);
	reader := bufio.NewReader(os.Stdin)
	game, _ := reader.ReadString('\n')
	fmt.Println(game)
	

}
