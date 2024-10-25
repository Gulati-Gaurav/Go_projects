package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// go fibre enables live updates in contrast to mux

// fake DB
var courses []Course

func main() {

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{FullName: "Hitesh Choudhary", Website: "go.dev"}})

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/allCourses", getAllCourses).Methods("GET")
	r.HandleFunc("/allCourse/{id}", getSingleCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/allCourse/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/allCourse/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers go in own separate file

func home(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1>Welcome Man</h1>"))
}

func getAllCourses(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json") // to set header

	// data, err := json.Marshal(courses)
	// CheckErrNil(err)
	// resp.Write(data)
	// or
	json.NewEncoder(resp).Encode(courses)
	// imp. See whenever such cases come where you first have to get byte[] from one and pass to another, there must be a easy way too.
}

// usually the functions are named as courses(for all), course(for single)
func getSingleCourse(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) // gives all variables in request
	fmt.Println(params)

	resp.Header().Set("Content-Type", "application/json")

	for _, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(resp).Encode(c)
			return
		}
	}
	json.NewEncoder(resp).Encode("No Course found")
}

func createCourse(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-Type", "application/json") // to set Header
	if req.Body == nil {
		json.NewEncoder(resp).Encode("Please send some data")
		return
	}
	var course Course
	json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(resp).Encode("Please send some data")
		return
	}
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(resp).Encode(course)
}

func updateCourse(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var c Course
			json.NewDecoder(req.Body).Decode(&c)
			c.CourseId = params["id"]
			courses = append(courses, c)
			json.NewEncoder(resp).Encode(c)
		}
	}
}

func deleteCourse(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}
	}
}

func CheckErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
