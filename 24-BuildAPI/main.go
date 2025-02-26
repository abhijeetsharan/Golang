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

// Model for courses - file
type Courses struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses = []Courses{}

// middleware, helper - file
func IsEmpty(c *Courses) bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building APIs")
	r := mux.NewRouter()

	// seeding 
	courses = append(courses, Courses{CourseId: "2", CourseName: "ReactJS", Price: 299, Author: &Author{FullName: "Abhijeet Sharan", Website: "lco.dev"}})
	courses = append(courses, Courses{CourseId: "4", CourseName: "MERN Stack", Price: 199, Author: &Author{FullName: "Abhijeet Sharan", Website: "go.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",  r))
}

// controllers- file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Courses
	_ = json.NewDecoder(r.Body).Decode(&course)
	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//Todo : check only if title is duplicate
	// loop, title matches with course.coursename, JSON
	// generate a unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// update course
			courses = append(courses[:index], courses[index+1:]...)

			var updatedCourse Courses
			_ = json.NewDecoder(r.Body).Decode(&course)
			updatedCourse.CourseId = params["id"]

			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	// TODO : send a response when id is not found
	json.NewEncoder(w).Encode("No course id found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	found := false

	// loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// delete course
			courses = append(courses[:index], courses[index+1:]...)
			found = true
			break
		}
	}
	if found {
		json.NewEncoder(w).Encode("Course deleted successfully")
	} else {
		json.NewEncoder(w).Encode("No course found")
	}
}