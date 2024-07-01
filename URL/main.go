package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	my_url := "https://jsonplaceholder.typicode.com/todos/1"
	parsedUrl, err := url.Parse(my_url)
	if(err!=nil){
		fmt.Println("Can't parse URL: ", err)
		return
	}
	fmt.Printf("Type of URL: %T", parsedUrl)
	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Domain Name: ", parsedUrl.Host)
	fmt.Println("Path: ", parsedUrl.Path)
	fmt.Println("RawQuery: ", parsedUrl.RawQuery)
}
