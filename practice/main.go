package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Problem struct{
	Name string `json:"name"`
}

type Member struct{
	Handle string `json:"handle"`
}

type Author struct{
	Members []Member `json:"members"`
}

type Cheater struct{
	CandidateId int `json:"id"`
	CandidateName Author `json:"author"`
	ContestId int `json:"contestId"`
	ProblemName Problem `json:"problem"`
	Verdict string `json:"verdict"`
}

type Apiresponse struct{
	Status string 
	Result []Cheater 
}

func main() {
	fmt.Println("CODEFORCES CONTEST STATUS...")
	getresponse()
}

func getresponse(){
	myurl := "https://codeforces.com/api/user.status?handle=Benq&from=1&count=20"

	response, err := http.Get(myurl)
	if(err!=nil){
		panic(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	
	var apiresponse Apiresponse
	err = json.Unmarshal(data, &apiresponse)
	if(err!=nil){
		panic(err)
	}
	fmt.Printf("Candidate Name: %s\n", apiresponse.Result[1].CandidateName.Members[0].Handle)

	for _, it := range apiresponse.Result{
		fmt.Printf("\nCandidateID: %d\n ContestID: %d\n Problem: %s\n Verdict: %s\n", it.CandidateId, it.ContestId, it.ProblemName.Name, it.Verdict)
	}

}