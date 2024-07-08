package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making server in backend...")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if(err!=nil){
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:3000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"price": 18,
			"platform": "youtube"
	}
	`)

	response , err := http.Post(myurl, "application/json", requestBody)
	if(err!=nil){
		panic(err)
	}

	defer response.Body.Close()

	content , _ := ioutil.ReadAll(response.Body)  //read all the response
	fmt.Println(string(content))

}