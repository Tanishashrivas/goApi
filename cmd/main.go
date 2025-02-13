package main

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/tanishashrivas/goApi/internal"
)

func main() {
	router := internal.SetupRoutes()

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
