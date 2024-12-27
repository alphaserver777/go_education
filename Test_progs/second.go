package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello"
	var d bool = true

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	sum := a + 5
	fmt.Println("Sum:", sum)

	isEqual := a == 10
	fmt.Println("Is equal:", isEqual)
}
