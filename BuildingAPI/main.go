package main

//encoder for sending data to the request body
//decoder for extracting data from the request

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course  -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helper -separate file
func (c *Course) IsEmpty() bool {
	return c.CourseName==""
}

func main() {
	fmt.Println("Building API...")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "Reactjs Bootcamp",
		CoursePrice: 399,
		Author: &Author{
			FullName: "Rekcon",
			Website: "www.rekcon.com",
		},
	})

	courses = append(courses, Course{
		CourseId: "5",
		CourseName: "Web3.0 Bootcamp",
		CoursePrice: 899,
		Author: &Author{
			FullName: "Hark",
			Website: "www.hark.com",
		},
	})

//routing
r.HandleFunc("/", serveHome).Methods("GET")
r.HandleFunc("/courses", getAllCourses).Methods("GET")
r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
r.HandleFunc("/course", CreateOneCourse).Methods("POST")
r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	//listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
	
}

//controllers -separate file

// serve home route i.e making a home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API Building</h1>"))
}

//one more controller/handler
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //whatever value is passed in encode, it will get treated as json value and will be thrown back to whoever making this request
}


func GetOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course...")
	w.Header().Set("Content-Type", "application/json")

	//grab course id from request
	params := mux.Vars(r)
	fmt.Println(params)

	//loop through courses, find matching id that user has sent, return the response
	for _, course := range courses{
		if(course.CourseId==params["id"]){  //id is given because in router there is name id given
			json.NewEncoder(w).Encode(course)
			return 
		}

	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}


func CreateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course...")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if(r.Body==nil){
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//what about data is send like this - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if(course.IsEmpty()){
		json.NewEncoder(w).Encode("No data inside JSON")
		return 
	}

	//generate unique id, convert into string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating one course...")
	w.Header().Set("Content-Type", "application/json")

	//first -grab id from request
	params := mux.Vars(r)

	//loop to find id, then remove, then add with my id(new id which is in params)
	for index, course := range courses{
		if(course.CourseId==params["id"]){
			courses = append(courses[:index], courses[index+1:]...) //removed the previous course
			var newcourse Course
			_ = json.NewDecoder(r.Body).Decode(&newcourse)
			newcourse.CourseId = params["id"]
			courses = append(courses, newcourse)
			json.NewEncoder(w).Encode(newcourse)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Delete One Course..")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop ,find id, remove(index, index+1)
	for index, course := range courses{
		if(course.CourseId==params["id"]){
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}
}