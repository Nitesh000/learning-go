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

// Model for course - separate file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware , helper funcitons - separate file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API HANDLING - with golang")
	fmt.Println("---------------------------")
	r := mux.NewRouter()

	// sedding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Nitesh Tudu", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Nitesh Tudu", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/course/delete-all}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - separate file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcoem to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// grabbing one course with the help of get api
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id frm request's param
	params := mux.Vars(r)

	// loop thorugh the courses
	// find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

// creating course (adding it to db){in this case it is inserting on the slice}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if recieved data is like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	// TODO: check only if title is duplicate
	// loop title matches ith course.coursename, JSON
	for _, c := range courses {
		if course.CourseName == c.CourseName {
			json.NewEncoder(w).Encode("Course already exist")
			return
		}
	}

	// genenrate unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update operation in golang
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop thorugh the values, id , remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			response := map[string]interface{}{
				"message": "This course has been updated",
				"course":  updatedCourse,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	// TODO: send a response when id is not found
	json.NewEncoder(w).Encode("Course not found with the similar id")
}

// delete operation in golan
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	// if the course not found then return something
	json.NewEncoder(w).Encode("No course found with the same id")
	return
}

// delete all course operation in { golan } this api

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all the courses")
	w.Header().Add("Content-Type", "application/json")

	// deleting all the slices
	courses = nil
	// another method is by using the append method
	// this is also quite great if we thing about it. we slice it by start to the zeroth element
	// courses = append(courses[:0])
	json.NewEncoder(w).Encode("All the courses data has been deleted")
	return
}
