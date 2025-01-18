package main

import (
	"errors"
	"fmt"
)

const pi = 3.14159265359

func main() {
	printCircleArea(-10)

}

func printCircleArea(circleRadius float32) {
	circleArea, err := calculatedCircleArea(circleRadius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Площадь круга с радиусом %.1f см: %.2f см²\n", circleRadius, circleArea)
	fmt.Print("Формула для расчёта площади круга Пr*2\n")

}

func calculatedCircleArea(circleRadius float32) (float32, error) {
	if circleRadius <= 0 {
		return float32(0), errors.New("Радиус отрицательный!")
	}
	return float32(circleRadius) * float32(circleRadius) * pi, nil
}
