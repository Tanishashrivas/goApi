package internal

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/tanishashrivas/goApi/internal/models"
	utils "github.com/tanishashrivas/goApi/pkg"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page for the APIs!"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range m.Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	utils.SendErrorResponse(w, "Course not found", http.StatusNotFound)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var newCourse m.Course
	if err := json.NewDecoder(r.Body).Decode(&newCourse); err != nil {
		utils.SendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateCourse(&newCourse); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// In Go, the math/rand package generates pseudo-random numbers, but by default, it produces the same sequence every time the program runs.
	// This happens because Go uses a fixed seed (by default, 1). A seed is like the starting point for random number generation.
	// rand.Seed(time.Now().UnixNano()) // time.Now().UnixNano() gets the current time in nanoseconds. Passing this to rand.Seed() makes sure every run starts with a different seed, so numbers appear truly random.
	// newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	rn, _ := rand.Int(rand.Reader, big.NewInt(100))
	newCourse.CourseId = rn.String()

	m.Courses = append(m.Courses, newCourse)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCourse)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	params := mux.Vars(r)
	var updatedCourse m.Course

	if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
		utils.SendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateCourse(&updatedCourse); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, course := range m.Courses {
		if course.CourseId == params["id"] {
			m.Courses[i] = updatedCourse
			m.Courses[i].CourseId = params["id"]
			json.NewEncoder(w).Encode(m.Courses[i])
			return
		}
	}

	utils.SendErrorResponse(w, "Course not found", http.StatusNotFound)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, course := range m.Courses {
		if course.CourseId == params["id"] {
			m.Courses = append(m.Courses[:i], m.Courses[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	utils.SendErrorResponse(w, "Course not found", http.StatusNotFound)
}
