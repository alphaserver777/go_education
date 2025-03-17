package main

import "fmt"

func main() {
	fmt.Println(Grow([]int{1, 2, 3, 4, 5}))
}

func Grow(nums []int) int {
	result := 1

	for _, num := range nums {
		result *= num
	}

	return result
}
