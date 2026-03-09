package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"student-api/internal/container"
)

func main() {
	app, err := container.New()
	if err != nil {
		log.Fatalf("Uygulama başlatılamadı: %v", err)
	}
	defer app.DB.Close()

	addr := ":" + strconv.Itoa(app.Config.Server.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Sistem ayakta!")
	})
	mux.HandleFunc("POST /students", app.StudentHandler.CreateStudent)
	mux.HandleFunc("GET /students", app.StudentHandler.ListStudents)
	mux.HandleFunc("GET /student", app.StudentHandler.GetStudent)
	mux.HandleFunc("DELETE /student", app.StudentHandler.DeleteStudent)

	fmt.Printf("Sunucu %s ortamında, %s portunda çalışıyor...\n", app.Config.Server.Env, addr)

	// 3. Server'ı başlat
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Sunucu hatası: %v", err)
	}
}
