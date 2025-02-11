package main

import (
	"fmt"
)

// 1. Объявляем интерфейс Number, который может быть либо int64, либо float64.
//   - ~int64 означает, что это может быть любой тип, основанный на int64 (включая, например, CustomInt).
//   - | - оператор объединения.
type Number interface {
	~int64 | float64
}

// 2. Создаем свой тип CustomInt, основанный на int64.
type CustomInt int64

// 3. Создаем метод IsPositive для типа CustomInt.
func (ci CustomInt) IsPositive() bool {
	return ci > 0
}

//  4. Определяем дженерик-тип Numbers, который представляет собой слайс элементов типа T,
//     где T должен реализовывать интерфейс Number.
type Numbers[T Number] []T

func main() {
	// Вызовы функций для демонстрации разных аспектов дженериков
	showSum()
	//showContains()
	//showAny()
	//unionInterfaceAndType()
	typeApproximation()
}

// 5. Функция showSum демонстрирует использование дженерик-функции sum.
func showSum() {
	// Создаем слайсы float64 и int64.
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}

	// Вызываем дженерик-функцию sum с разными типами данных.
	fmt.Println(sum(floats))      // Автоматическое определение типа float64.
	fmt.Println(sum[int64](ints)) // Явное указание типа int64.

	// так как слайс []interface{} не подходит под ограничения дженерик-функции sum.
	// wrongFloats := []interface{}{"string", struct{}{}, true}
	// fmt.Println(sum(wrongFloats))
}

// 6. Функция showContains демонстрирует использование дженерик-функции contains.
func showContains() {
	// Объявляем структуру Person.
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	// Создаем слайсы разных типов.
	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", contains(ints, 4)) // Проверяем наличие 4 в слайсе ints.

	strings := []string{"Vasya", "Dima", "Katya"}
	fmt.Println("strings:", contains(strings, "Katya")) // Проверяем наличие "Katya" в слайсе strings.
	fmt.Println("strings:", contains(strings, "Sasha")) // Проверяем наличие "Sasha" в слайсе strings.

	people := []Person{ // Создаем слайс структур Person.
		{
			name:     "Vasya",
			age:      20,
			jobTitle: "Programmer",
		},
		{
			name:     "Dasha",
			age:      23,
			jobTitle: "Designer",
		},
		{
			name:     "Pasha",
			age:      30,
			jobTitle: "Admin",
		},
	}

	// Проверяем наличие структур в слайсе people.
	fmt.Println("structs:", contains(people, Person{
		name:     "Vasya",
		age:      21,
		jobTitle: "Programmer",
	}))

	fmt.Println("structs:", contains(people, Person{
		name:     "Vasya",
		age:      20,
		jobTitle: "Programmer",
	}))
}

// 7. Функция showAny демонстрирует использование дженерик-функции show с типом any.
func showAny() {
	// Вызываем дженерик-функцию show с разными типами данных.
	show(1, 2, 3)                            // Слайс int
	show("test1", "test2", "test3")          // Слайс string
	show([]int64{1, 2, 3}, []int64{4, 5, 6}) // Слайс слайсов int64
	show(map[string]int64{                   // Map
		"first":  1,
		"second": 2,
	})
	//  Приведение типов к интерфейсу.
	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "Vasya"}))
}

// 8. Функция unionInterfaceAndType демонстрирует использование дженерик-типа Numbers.
func unionInterfaceAndType() {
	// Создаем переменные типа Numbers с разными типами данных.
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...) // Добавляем значения в слайс.

	floats := Numbers[float64]{1.0, 2, 5, 3, 5} // Создаем слайс и инициализируем его значениями.

	// Вызываем дженерик-функцию sumUnionInterface с разными типами данных.
	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

// 9. Функция typeApproximation демонстрирует использование дженерик-функции с CustomInt.
func typeApproximation() {
	// Создаем слайс типа CustomInt.
	customInts := []CustomInt{1, 2, 3, 5, 6}

	// Создаем слайс типа int64 и конвертируем в него слайс CustomInt
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		castedInts[idx] = int64(val)
	}

	// Вызываем дженерик-функцию sumUnionInterface с типами CustomInt и int64.
	fmt.Println(sumUnionInterface(customInts))
	fmt.Println(sumUnionInterface(castedInts))
}

// 10. Дженерик-функция contains проверяет наличие элемента в слайсе.
func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements { // Проходим по слайсу
		if searchEl == el { // Если значение найдено, возвращаем true
			return true
		}
	}
	return false // Если не нашли, возвращаем false
}

// 11. Дженерик-функция sum вычисляет сумму элементов слайса.
func sum[V int64 | float64](numbers []V) V { // constrain - ограничение типов Дженерика
	var sum V // Объявляем переменную sum типа V, которая будет хранить результат
	for _, num := range numbers {
		sum += num // Прибавляем к sum значение из слайса
	}
	return sum // Возвращаем результат
}

// 12. Дженерик-функция sumUnionInterface вычисляет сумму элементов слайса с типом, который реализует интерфейс Number.
func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// 13. Дженерик-функция show выводит на экран слайс любого типа
func show[T any](entities ...T) {
	fmt.Println(entities)
}
