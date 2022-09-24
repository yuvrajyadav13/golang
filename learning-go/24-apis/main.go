package main

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

// Model for courses
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Create a slice for storing courses
var courses []Course

// middleware or helper functions

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building APIs in GO")

	r := mux.NewRouter()

	// Seeding: adding fake data
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", Price: 299, Author: &Author{Fullname: "Gola", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "MERN", Price: 199, Author: &Author{Fullname: "Gola1", Website: "gola.dev"}})

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controllers or the functions that handles requests

// Serve Home

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to APIs in GO<h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// Loop through courses and find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about {} or []
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("CourseName is not provided provided")
		return
	}

	// generate unique id, string and append course into courses slice

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one of the course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id
	params := mux.Vars(r)

	// loop, search for ID.
	// Remove the course
	// Update the course and append it again

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// If courseid is not found
	json.NewEncoder(w).Encode("Course ID not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is removed")
			return
		}
	}

	// If courseId is not found
	json.NewEncoder(w).Encode("Course ID not found")
	return
}
