package database

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string, migrationsPath string) error {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("migration başlatılamadı: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("Veritabanı zaten güncel, yeni migration yok.")
			return nil
		}
		return fmt.Errorf("migration uygulanırken hata: %w", err)
	}

	slog.Info("Migration işlemleri başarıyla tamamlandı.")
	return nil
}
