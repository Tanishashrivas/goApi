package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "github.com/tanishashrivas/goApi/internal/models"
)

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ValidateCourse(course *m.Course) error {
	if course.CourseName == "" {
		return fmt.Errorf("course name is required")
	}
	if course.CoursePrice <= 0 {
		return fmt.Errorf("course price must be greater than zero")
	}
	if course.Author == nil || course.Author.FullName == "" {
		return fmt.Errorf("author details are required")
	}
	return nil
}

func CheckNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
