package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/tanishashrivas/goApi/internal/models"
	utils "github.com/tanishashrivas/goApi/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home page for the APIs!")
	w.Write([]byte("Home page for the APIs!"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var courses []m.Course

	cursor, err := Collection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())

	if err != nil {
		utils.SendErrorResponse(w, "Error finding courses", http.StatusInternalServerError)
		return
	}

	if err := cursor.All(context.TODO(), &courses); err != nil {
		utils.SendErrorResponse(w, "Error decoding courses", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.SendErrorResponse(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	var course m.Course

	err = Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&course)

	if err != nil {
		utils.SendErrorResponse(w, "Course not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(course)
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

	// rn, _ := rand.Int(rand.Reader, big.NewInt(100))
	// newCourse.CourseId = rn.String()

	// m.Courses = append(m.Courses, newCourse)
	inserted, err := Collection.InsertOne(context.Background(), newCourse)

	utils.CheckNilError(err)

	fmt.Println("Ïnserted one course in DB", inserted)

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

	id, err := primitive.ObjectIDFromHex(params["id"])

	utils.CheckNilError(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedCourse}

	result, err := Collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		utils.SendErrorResponse(w, "Error updating course", http.StatusInternalServerError)
		return
	}

	if result.ModifiedCount == 0 {
		utils.SendErrorResponse(w, "Course not found or not modified", http.StatusNotFound)
		return
	}

	// for i, course := range m.Courses {
	// 	if course.CourseId == params["id"] {
	// 		m.Courses[i] = updatedCourse
	// 		m.Courses[i].CourseId = params["id"]
	// 		json.NewEncoder(w).Encode(m.Courses[i])
	// 		return
	// 	}
	// }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCourse)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendErrorResponse(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	result, err := Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		utils.SendErrorResponse(w, "Error deleting course", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		utils.SendErrorResponse(w, "Course not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	// for i, course := range m.Courses {
	// 	if course.CourseId == params["id"] {
	// 		m.Courses = append(m.Courses[:i], m.Courses[i+1:]...)
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	// }
}
