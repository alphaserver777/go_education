package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver. Change if needed.
)

type Car struct {
	ID       uuid.UUID `json:"id"`
	Order    int       `json:"order"`
	FullName string    `json:"fullName"`
	CarBrand string    `json:"carBrand"`
	Plate    string    `json:"plate"`
	Note     string    `json:"note"`
}

func main() {
<<<<<<< HEAD
	db, err := sql.Open("postgres", "user=postgres password=K#7sd4Na dbname=postgres sslmode=disable")
=======
	db, err := sql.Open("postgres", "user=postgres password=K#7sd4Na dbname=your_db sslmode=disable") // Replace with your connection string
>>>>>>> 38011a1dc06b4ada43ffaed2cb442808b3ea67a8
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Ошибка закрытия соединения с базой данных: %v", err)
		}
	}()

	// Create table if not exists (ensure this is executed only once during deployment)
	_, err = db.Exec(`
  CREATE TABLE IF NOT EXISTS cars (
   id UUID PRIMARY KEY,
   order_num INT,
   full_name TEXT,
   car_brand TEXT,
   plate TEXT UNIQUE,
   note TEXT
  )
 `)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/car/{plate}", getCarHandler(db)).Methods("GET")
	router.HandleFunc("/car", addCarHandler(db)).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func getCarHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		platePart := vars["plate"]

		rows, err := db.Query("SELECT id, order_num, full_name, car_brand, plate, note FROM cars WHERE plate LIKE '%' || $1 || '%'", platePart)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка выполнения запроса: %v", err), http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := rows.Close(); err != nil {
				log.Printf("Ошибка закрытия результата запроса: %v", err)
			}
		}()

		cars := []Car{}
		for rows.Next() {
			var car Car
			err := rows.Scan(&car.ID, &car.Order, &car.FullName, &car.CarBrand, &car.Plate, &car.Note)
			if err != nil {
				http.Error(w, fmt.Sprintf("Ошибка чтения данных: %v", err), http.StatusInternalServerError)
				return
			}
			cars = append(cars, car)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка итерации строк: %v", err), http.StatusInternalServerError)
			return
		}

		if len(cars) == 0 {
			http.Error(w, "Машин не найдено", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	}
}

func addCarHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// ... (addCarHandler function remains largely the same) ...
	}
}
