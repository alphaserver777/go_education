package main

import (
	"fmt"

	"github.com/google/uuid"
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
	newCar := Car{
		ID:       uuid.New(),
		Order:    1,
		FullName: "Иван Иванов",
		CarBrand: "Toyota",
		Plate:    "A123BB",
		Note:     "Хорошая машина",
	}

	fmt.Println(newCar)
}
