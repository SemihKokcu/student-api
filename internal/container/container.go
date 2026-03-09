package container

import (
	"database/sql"
	"log/slog"
	"student-api/internal/config"
	"student-api/internal/database"
	"student-api/internal/handlers"
	"student-api/internal/logger"
	"student-api/internal/repository"
)

type Container struct {
	Config         *config.Config
	DB             *sql.DB
	Logger         *slog.Logger
	StudentHandler *handlers.StudentHandler
}

func New() (*Container, error) {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	//  create logger
	l := logger.SetupLogger(cfg.Log.Format, cfg.Log.Level)
	l.Info("Uygulama yapılandırması yüklendi", "env", cfg.Server.Env)

	// connect to database
	db, err := database.ConnectDB(
		cfg.Database.URL,
		cfg.Database.MaxOpenConns,
		cfg.Database.MaxIdleConns,
	)
	if err != nil {
		l.Error("Veritabanı bağlantı hatası", "hata", err)
		return nil, err
	}
	l.Info("Veritabanına başarıyla bağlanıldı")

	// create injections
	studentRepo := repository.NewStudentRepository(db)
	studentHandler := handlers.NewStudentHandler(studentRepo)

	// return container
	return &Container{
		Config:         cfg,
		DB:             db,
		Logger:         l,
		StudentHandler: studentHandler,
	}, nil
}
