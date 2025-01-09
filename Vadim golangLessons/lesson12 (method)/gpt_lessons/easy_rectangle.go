package main

import "fmt"

type Rectangle struct { // Объявил структура Треугольника со значениями Ширины и Высоты
	Wight  float64
	Height float64
}

func (r Rectangle) Area() float64 { // Объявление метода Площади, который работает со структурой прямоуголька
	return r.Height * r.Wight
}

func main() {
	fmt.Println(Rectangle{5, 20}.Area())
}
