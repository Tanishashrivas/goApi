package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/tanishashrivas/goApi/internal/models"
)

func ServeHome(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Home page for the APIs!"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for _, course := range m.Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given course id")
}

// func CreateCourse(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	m.Courses = append(m.Courses, )
// 	json.NewEncoder(w).Encode([]byte("No course with that course id found"))
// }
