package container

import (
	"database/sql"
	"student-api/internal/config"
	"student-api/internal/database"
	"student-api/internal/handlers"
	"student-api/internal/repository"
)

type Container struct {
	Config         *config.Config
	DB             *sql.DB
	StudentHandler *handlers.StudentHandler
}

func New() (*Container, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.ConnectDB(
		cfg.Database.URL,
		cfg.Database.MaxOpenConns,
		cfg.Database.MaxIdleConns,
	)
	if err != nil {
		return nil, err
	}

	studentRepo := repository.NewStudentRepository(db)
	studentHandler := handlers.NewStudentHandler(studentRepo)

	return &Container{
		Config:         cfg,
		DB:             db,
		StudentHandler: studentHandler,
	}, nil
}
