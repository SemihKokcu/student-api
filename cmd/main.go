package main

import (
	"fmt"
	"log"
	"net/http"
	"student-api/internal/database"
	"student-api/internal/handlers"
	"student-api/internal/repository"
)

func main() {
	// 1. Veritabanı bağlantısını kur
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer db.Close()

	// 2. Repository'yi oluştur (Veritabanını içine enjekte et)
	studentRepo := repository.NewStudentRepository(db)

	// 3. Handler'ı oluştur (Repository'yi içine enjekte et)
	// Not: Handler artık repository'ye ihtiyaç duyuyor
	studentHandler := handlers.NewStudentHandler(studentRepo)

	mux := http.NewServeMux()

	// 4. Rotaları bağla
	mux.HandleFunc("POST /students", studentHandler.CreateStudent)
	mux.HandleFunc("GET /students", studentHandler.ListStudents)
	mux.HandleFunc("GET /student", studentHandler.GetStudent)
	mux.HandleFunc("DELETE /student", studentHandler.DeleteStudent)

	fmt.Println("Server :8080 portunda ve PostgreSQL ile çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
