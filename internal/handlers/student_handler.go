package handlers

import (
	"strconv"

	"student-api/internal/models"
	"student-api/internal/repository"
	"student-api/internal/response"
	"student-api/internal/validator"

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
		response.ValidationError(c, validator.MapValidationErrors(err))
		return
	}

	if err := h.repo.Create(s); err != nil {
		response.InternalError(c, "Öğrenci kaydedilemedi.")
		return
	}

	response.Created(c, s)
}

func (h *StudentHandler) ListStudents(c *gin.Context) {
	students, err := h.repo.GetAll()
	if err != nil {
		response.InternalError(c, "Öğrenciler yüklenemedi.")
		return
	}
	response.Success(c, students)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Geçersiz ID.")
		return
	}

	student, err := h.repo.GetByID(id)
	if err != nil {
		response.NotFound(c, "Öğrenci bulunamadı.")
		return
	}

	response.Success(c, student)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Geçersiz ID.")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		response.InternalError(c, "Öğrenci silinemedi.")
		return
	}

	response.Success(c, nil)
}
