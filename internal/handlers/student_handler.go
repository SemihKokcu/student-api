package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-api/internal/models"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	s.ID = models.NextID
	students, err := models.LoadFromFile()
	if err != nil {
		http.Error(w, "Öğrenci yüklenemedi", http.StatusInternalServerError)
		return
	}
	students = append(students, s)
	result := models.SaveToFile(students)
	if result != nil {
		http.Error(w, "Öğrenci kaydedilemedi", http.StatusInternalServerError)
		return
	}
	models.NextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func ListStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	students, err := models.LoadFromFile()
	if err != nil {
		http.Error(w, "Öğrenci yüklenemedi", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}
	students, err := models.LoadFromFile()
	if err != nil {
		http.Error(w, "Öğrenci yüklenemedi", http.StatusInternalServerError)
		return
	}
	for _, s := range students {
		if s.ID == id {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "Öğrenci bulunamadı", http.StatusNotFound)
}
