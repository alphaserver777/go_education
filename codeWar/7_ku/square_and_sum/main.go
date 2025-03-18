/*
Complete the square sum function so that it squares each number
 passed into it and then sums the results together.

For example, for [1, 2, 2] it should return 9
*/

package main

func main() {
	SquareSum([]int{1, 2, 2})

}

func SquareSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}
