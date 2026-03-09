package repository

import (
	"database/sql"
	"student-api/internal/models"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(s models.Student) error {
	query := `INSERT INTO students (name, grade) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, s.Name, s.Grade).Scan(&s.ID)
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.db.Query("SELECT id, name, grade FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Grade); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func (r *StudentRepository) GetByID(id int) (*models.Student, error) {
	var s models.Student
	query := `SELECT id, name, grade FROM students WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&s.ID, &s.Name, &s.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) Delete(id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
