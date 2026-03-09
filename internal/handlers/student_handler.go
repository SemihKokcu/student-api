package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-api/internal/models"
	"student-api/internal/repository"
)

type StudentHandler struct {
	repo *repository.StudentRepository
}

func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(s); err != nil {
		http.Error(w, "Öğrenci oluşturulamadı", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (h *StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	students, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Öğrenci yüklenemedi", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}
	students, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Öğrenci yüklenemedi", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Öğrenci silinemedi", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
