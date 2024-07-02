package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	// we want to take only these three data, so we make this struct
	UserID    int    `json: "userId"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

//Get method
func performGetRequest(){
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

//Post Method
func performPostRequest()  {
	todo := Todo {
		UserID: 22,
		Title: "Using post method",
		Completed: true,
	}

	//convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if(err!=nil){
		fmt.Println("Error while marshalling: ", err)
		return
	}

	//convert the json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	//send post request
	response , err := http.Post(myUrl, "application/json", jsonReader)
	if(err!=nil){
		fmt.Println("Error sending request: ", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

func main() {
	fmt.Println("Learning CRUD...")
	performGetRequest()
	performPostRequest()

}
