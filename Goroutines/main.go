package main

import (
	"fmt"
	"time"
)

func sayHello()  {
	fmt.Println("Hello, world!")
	time.Sleep(2000 * time.Millisecond)
}

func sayHi(){
	fmt.Println("Say Hi :-)")
}

func main() {
	fmt.Println("Learning goroutine...")
	go sayHello()
	sayHi()

	time.Sleep(200*time.Millisecond)
}