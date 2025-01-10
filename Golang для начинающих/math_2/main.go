package main

import (
	"bufio"   // Пакет для работы с буферизированным вводом-выводом
	"fmt"     // Пакет для форматированного ввода-вывода
	"math"    // Пакет для математических функций
	"os"      // Пакет для работы с операционной системой
	"strconv" // Пакет для преобразования строк в числа
	"strings" // Пакет для работы со строками
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Создаем объект для чтения данных с консоли

	fmt.Println("Выберите фигуру:") // Выводим меню выбора фигуры
	fmt.Println("1. Круг")
	fmt.Println("2. Квадрат")
	fmt.Println("3. Трапеция")

	figureChoice, err := readInt(reader, "Введите номер фигуры (1-3): ") // Читаем выбор пользователя и проверяем на ошибки
	if err != nil {
		fmt.Println("Ошибка ввода:", err) // Выводим сообщение об ошибке, если ввод некорректен
		return
	}

	switch figureChoice { // Выполняем вычисления в зависимости от выбора фигуры
	case 1:
		radius, err := readFloat(reader, "Введите радиус круга: ") // Читаем радиус круга и проверяем на ошибки
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		area := calculateCircleArea(radius)       // Вычисляем площадь круга
		fmt.Printf("Площадь круга: %.2f\n", area) // Выводим результат
	case 2:
		side, err := readFloat(reader, "Введите сторону квадрата: ") // Читаем сторону квадрата и проверяем на ошибки
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		area := calculateSquareArea(side)            // Вычисляем площадь квадрата
		fmt.Printf("Площадь квадрата: %.2f\n", area) // Выводим результат
	case 3:
		base1, err := readFloat(reader, "Введите длину первого основания трапеции: ") // Читаем первое основание трапеции и проверяем на ошибки
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		base2, err := readFloat(reader, "Введите длину второго основания трапеции: ") // Читаем второе основание трапеции и проверяем на ошибки
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		height, err := readFloat(reader, "Введите высоту трапеции: ") // Читаем высоту трапеции и проверяем на ошибки
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		area := calculateTrapezoidArea(base1, base2, height) // Вычисляем площадь трапеции
		fmt.Printf("Площадь трапеции: %.2f\n", area)         // Выводим результат
	default:
		fmt.Println("Неверный номер фигуры.") // Выводим сообщение, если введен неверный номер фигуры
	}
}

func calculateCircleArea(radius float64) float64 { // Функция для вычисления площади круга
	return math.Pi * radius * radius
}

func calculateSquareArea(side float64) float64 { // Функция для вычисления площади квадрата
	return side * side
}

func calculateTrapezoidArea(base1, base2, height float64) float64 { // Функция для вычисления площади трапеции
	return (base1 + base2) * height / 2
}

func readInt(reader *bufio.Reader, prompt string) (int, error) { // Функция для чтения целого числа с консоли с обработкой ошибок
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("некорректный ввод числа: %w", err)
	}
	return num, nil
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) { // Функция для чтения вещественного числа с консоли с обработкой ошибок
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("некорректный ввод числа: %w", err)
	}
	return num, nil
}
