package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if(err!=nil){
		fmt.Println("Error while creating file: ", err)
		return
	}
	
	//inserting data in a file
	content := "Hello world from world"
	_, errors := io.WriteString(file, content)
	if(errors != nil){
		fmt.Println("Error while writing in the file: ", errors)
		return
	}
	defer file.Close()

}