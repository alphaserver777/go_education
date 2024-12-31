//https://www.youtube.com/watch?v=YhUAHrovsH8&t=30s

package main

import (
	"fmt"
)

// Создаем структуру Human c полям: Name, Surname, Age -> По структуре создаем объект (Вадим) -> К объекту применяем метод SayHello().
type Human struct {
	Name    string // Поле структуры
	Surname string // Поле структуры
	Age     int    // Поле структуры
}

// Создаем метод SayHello() для структуры Human
func (h Human) SayHello() {
	fmt.Printf("Hello, my name is %s %s. I'm %d years old.\n", h.Name, h.Surname, h.Age)
}

func main() {
	// Создаем объект типа Human
	Vadim := Human{
		Name:    "Vadim",
		Surname: "Moiseenko",
		Age:     26,
	}
	// Создаем объект типа Human
	Nika := Human{
		Name:    "Nika",
		Surname: "Moiseenko",
		Age:     24,
	}

	fmt.Println(Vadim)

	fmt.Println(Nika)

	// Вызываем метод SayHello() для объекта Nikita
	Nika.SayHello()
	// Вызываем метод SayHello() для объекта Vadim
	Vadim.SayHello()
}
