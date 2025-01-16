package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"math/rand"

	"github.com/gorilla/mux"
)

// crud operation of courses database : basically it is a fake database creating
var courses []Course

// Model for coursese
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"` // here Author type is pointer of struct
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// middleware or helper mathod which checks some tasks
func (c *Course) IsEmpty() bool {
	return c.CourseName == "" // return true is there is no coureseid and coursename
}

func main() {
	fmt.Println("Welcome to api!!!!")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Gern Stack", CoursePrice: 2345, Author: &Author{
		Fullname: "Mirazul Alalm", Website: "goo.us"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "Hern Stack", CoursePrice: 3456, Author: &Author{
		Fullname: "siraz", Website: "icrew.us"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "pern Stack", CoursePrice: 564574, Author: &Author{
		Fullname: "Kairul Alam", Website: "go.dev"}})

	courses = append(courses, Course{CourseId: "5", CourseName: "mern Stack", CoursePrice: 2346745, Author: &Author{
		Fullname: "Arfat Ahmed", Website: "agiledot.us"}})

	// routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	// Listen to a port
	fmt.Println("Listening to port 8080 ")
	log.Fatal(http.ListenAndServe(":8080", r))

}

//controllers : serve the route

// Serve home router

func ServeHome(w http.ResponseWriter, r *http.Request) { // r will read the request and w will write as the respose of this request
	w.Write([]byte("<h1>Welcome to the Api by Miraz</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses!!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find maching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id!!")
	return
}

// create course or adding course to db
func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One course!")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("The course name is null")
		return
	}

	// todo : check only if the title is duplicate
	// loop , title match with course.CourseName , json

	// Create Unique id
	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// append course into Courses(db)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course!!")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// first - get id from request
	params := mux.Vars(r)

	// loop id , remove , and add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // removing that value
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found")
	return
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One course!!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop , id , remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // this will delete the value of type Struct from the courses db
			break
		}
	}
}
