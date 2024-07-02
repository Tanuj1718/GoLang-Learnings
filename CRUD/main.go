package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	// we want to take only these three data, so we make this struct
	UserID    int    `json: "userId"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

func main() {
	fmt.Println("Learning CRUD...")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while getting response: ", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error while getting the data: ", response.Status)
		return
	}

	// data, err := ioutil.ReadAll(response.Body)
	// if(err!=nil){
	// 	fmt.Println("Error while getting the data: ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	//api se hme kuch specific information hi chahiye toh hm isliye struct bnate h
	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo) //newdecoder se decode kiya todo ko
	if(err!=nil){
		fmt.Println("Error")
		return
	}

	fmt.Println(todo)

}
