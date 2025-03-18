/*
Given an array of integers your solution should find the smallest integer.

For example:

    Given [34, 15, 88, 2] your solution will return 2
    Given [34, -345, -1, 100] your solution will return -345

You can assume, for the purpose of this kata, that the supplied array will not be empty.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers1 := []int{34, 15, 88, 2}
	smallest1 := FindSmallest3(numbers1)
	fmt.Println("Smallest in [34, 15, 88, 2]:", smallest1) // Output: Smallest in [34, 15, 88, 2]: 2

	numbers2 := []int{34, -345, -1, 100}
	smallest2 := FindSmallest3(numbers2)
	fmt.Println("Smallest in [34, -345, -1, 100]:", smallest2) // Output: Smallest in [34, -345, -1, 100]: -345
}

func FindSmallest1(numbers []int) int {
	smallest := numbers[0]

	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

func FindSmallest2(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func FindSmallest3(numbers []int) int {
	smallest := numbers[0]
	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
