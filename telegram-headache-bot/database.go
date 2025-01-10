package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

// DB - это глобальная переменная для работы с БД
var DB *sql.DB

func initDB(databaseURL string) {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Создание таблицы, если ее нет
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS headache_reports (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL,
            date DATE NOT NULL DEFAULT CURRENT_DATE,
            answer TEXT NOT NULL
        );
    `
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Database connection established and table ready.")
}

func SaveReport(userID int64, answer string) error {
	_, err := DB.Exec("INSERT INTO headache_reports (user_id, answer) VALUES ($1, $2)", userID, answer)
	return err
}
