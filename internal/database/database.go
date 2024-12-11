package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/christmas-fire/Golang_PostgreSQL_Init/internal/config"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	cfg, err := config.LoadConfig("./internal/config/")
	if err != nil {
		log.Fatal(err)
	}

	con := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port, cfg.Sslmode,
	)

	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("БД недоступна: %v", err)
	}

	log.Println("Подключение к БД прошло успешно")

	return db
}

func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			role TEXT NOT NULL
		)`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу: %w", err)
	}

	log.Println("Таблица создана успешно")
	return nil
}
