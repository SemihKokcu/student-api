package main

import (
	"fmt"
	"net/http"
	"student-api/internal/handlers"
	"student-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Rotalar
	mux.HandleFunc("POST /students", handlers.CreateStudent)
	mux.HandleFunc("GET /students", handlers.ListStudents)
	mux.HandleFunc("GET /student", handlers.GetStudent)

	// Middleware ile sarmalla
	wrappedMux := middleware.RequestLogger(mux)

	fmt.Println("Öğrenci Sistemi 8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", wrappedMux)
}
