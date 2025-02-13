package internal

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/course", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/course", CreateCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")

	return r
}
