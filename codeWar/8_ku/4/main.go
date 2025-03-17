/*
Nathan loves cycling.

Because Nathan knows it is important to stay hydrated, he drinks 0.5 litres of water per hour of cycling.

You get given the time in hours and you need to return the number of litres Nathan will drink, rounded to the smallest value.

For example:

time = 3 ----> litres = 1

time = 6.7---> litres = 3

time = 11.8--> litres = 5
*/
package main

import "fmt"

func main() {
	fmt.Println(Litres(3))
	data := []float64{3.0, 6.7, 11.8}
	fmt.Println(kmLitres(data))
}

func Litres(time float64) int {
	return int(time * 0.5)
}

func kmLitres(km []float64) []int {
	var result []int

	for _, i := range km {
		result = append(result, int(i*0.5))
	}
	return result
}
