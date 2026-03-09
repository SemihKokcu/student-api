package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Sürücüyü kaydet
)

func ConnectDB() (*sql.DB, error) {
	dsn := "postgres://admin:secret@localhost:5432/studentdb?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Bağlantıyı test et
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("PostgreSQL bağlantısı başarılı!")
	return db, nil
}
