package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if(err!=nil){
		fmt.Println("Error while getting response: ", err)
		return
	}
	defer response.Body.Close() //it is used to ensure that the response body is closed after you have finished reading from it. It's important to close resources like network connections and file handles to free up system resources

	data, err := ioutil.ReadAll(response.Body)
	if(err!=nil){
		fmt.Println("Error while extracting data: ", err)
		return
	}

	fmt.Println(string(data)) //data will be returned in bytes through api and we have to convert it into string
}