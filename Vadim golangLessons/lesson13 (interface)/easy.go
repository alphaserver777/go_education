package main

import "fmt"

// Определяем интерфейс с общим поведением
type Animal interface {
	Speak() string // Метод Speak, который возвращает строку
}

// Реализуем типы, которые соответствуют этому интерфейсу
type Dog struct {
	Name  string
	Breed string
}
type Cat struct {
	Name  string
	Breed string
}
type Bird struct {
	Name string
}

// Speak returns the sound a Dog makes, which is "Гав".
func (d Dog) Speak() string {
	return fmt.Sprintf("%s say Gav! She is %s", d.Name, d.Breed)
}

// Speak returns the sound a Cat makes, which is "Мяу".
func (c Cat) Speak() string {
	return fmt.Sprintf("%s say Mewow! She is %s", c.Name, c.Breed)
}

func (b Bird) Speak() string {
	return fmt.Sprintf("%s say Chirik!", b.Name)
}

// Функция, принимающая любой тип, соответствующий интерфейсу Animal
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	// Создаем разные объекты
	dog := Dog{Name: "Bonya", Breed: "Corgi"} //Создается экземпляр структуры Dog и присваивается переменной dog.
	cat := Cat{Name: "Jesica", Breed: "Siam"} //Создается экземпляр структуры Cat и присваивается переменной cat.
	bird := Bird{Name: "Ara"}

	// Вызываем одну и ту же функцию для разных типов
	MakeSound(dog)
	MakeSound(cat)
	MakeSound(bird)
}
