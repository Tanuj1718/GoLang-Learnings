package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CF struct{
	Id int `json:"contestId"`
	ContestName string `json:"contestName"`
	Rank int `json:"rank"`
}

type ApiResponse struct{
	Status string
	Result []CF
}

func main() {
	fmt.Println("Api from web...")
	getresponse()
}

func getresponse() {
	myurl := "https://codeforces.com/api/user.rating?handle=eyewal_18"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err:=ioutil.ReadAll(response.Body)
	if(err!=nil){
		panic(err)
	}

	fmt.Println(string(data))

	var apiresponse ApiResponse
	err = json.Unmarshal(data, &apiresponse)
	if(err!=nil){
		panic(err)
	}

	for _, a := range apiresponse.Result{
		fmt.Printf("CONTEST ID: %d\n CONTEST NAME: %s\n RANK: %d\n", a.Id, a.ContestName, a.Rank)
	}
}
