package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` //'-' will remove this field for whoever will consume api
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json handling ...")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 599, "LearnCodeOnline.in", "abcd123", []string{"web-dev", "js"}},
		{"Android Bootcamp", 599, "LearnCodeOnline.in", "abce123", []string{"web-dev", "js"}},
		{"DSA Bootcamp", 299, "LearnCodeOnline.in", "abcf123", nil},
	}

	//package this data as json data
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)

}


func DecodeJson(){
	jsonDataFromWeb := []byte(`
	        {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",   
                "tags": ["web-dev","js"],
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if(checkValid){
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}
}
