package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json: "name"`
	Age     int    `json: "age"`
	IsAdult bool   `json: "is_Adult"` //key used by json after converting object into json
}

func main() {
	fmt.Println("Learning json handling")
	person:= Person{Name: "Garima", Age: 20, IsAdult: true}
	fmt.Println(person)

	//convert person into json (Encoding/Marshalling)
	jsonData, err := json.Marshal(person)
	if(err!=nil){
		fmt.Println("Error while marshalling: ", err)
		return
	}
	fmt.Println(string(jsonData))
}
