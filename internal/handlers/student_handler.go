package handlers

import (
	"net/http"
	"strconv"

	"student-api/internal/models"
	"student-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	repo *repository.StudentRepository
}

func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var s models.Student

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	if err := h.repo.Create(s); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kaydedilemedi"})
		return
	}

	c.JSON(http.StatusCreated, s)
}

func (h *StudentHandler) ListStudents(c *gin.Context) {
	students, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Öğrenci yüklenemedi"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	student, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Öğrenci bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}
	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Öğrenci silinemedi"})
		return
	}
	c.Status(http.StatusNoContent)
}
