/*
How it`s work:
## 🔥 **Как это работает?**

1. **Мы определили интерфейс `Developer`**, который требует два метода: `Code()` и `Review()`.
2. **`JuniorDev` реализует только `Code()`, но не `Review()`**, поэтому он **не подходит**.
3. **`SeniorDev` реализует оба метода**, поэтому он проходит "собеседование".
4. **Функция `Hire(dev Developer)` принимает только тех, кто реализует весь интерфейс**.
5. Когда мы вызываем `Hire(junior)`, код не компилируется, потому что `JuniorDev` **не реализует** `Review()`.
*/
package main

import "fmt"

// Интерфейс = Вакансия (Требования к работнику)
type Developer interface {
	Code()   // Умение писать код
	Review() // Умение делать код-ревью
}

// имплементация интерфейса означает, что структура реализует (выполняет) все методы,
// которые требует интерфейс.

// Кандидат 1: Junior-разработчик (не проходит, потому что не делает Review)
type JuniorDev struct{}

func (j JuniorDev) Code() {
	fmt.Println("Junior пишет код...")
}

// JuniorDev НЕ имплементирует интерфейс Developer,
// потому что **нет метода Review()** ❌

// Кандидат 2: Senior-разработчик (подходит, потому что умеет и кодить, и делать ревью)
type SeniorDev struct{}

func (s SeniorDev) Code() {
	fmt.Println("Senior пишет код...")
}

func (s SeniorDev) Review() {
	fmt.Println("Senior делает код-ревью...")
}

// Компилятор проверит:
// - Есть ли у SeniorDev метод Code()? ✅ Да
// - Есть ли у SeniorDev метод Review()? ✅ Да
// → Значит, SeniorDev **имплементирует** интерфейс Developer 🎉

// Функция "приема на работу", которая принимает только тех, кто соответствует интерфейсу (вакансии)
func Hire(dev Developer) {
	fmt.Println("На собеседовании...")
	dev.Code()
	dev.Review()
	fmt.Println("Вы нам подходите!\n")
}

func main() {
	// junior := JuniorDev{}
	senior := SeniorDev{}

	fmt.Println("Пробуем нанять JuniorDev:")
	fmt.Println("❌ Ошибка! JuniorDev не реализует метод Review()")
	// Hire(junior) // ❌ Ошибка! JuniorDev не реализует метод Review()

	fmt.Println("\nПробуем нанять SeniorDev:")
	Hire(senior) // ✅ Успешно! SeniorDev полностью соответствует требованиям
}
