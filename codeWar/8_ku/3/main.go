/*
Write a function that accepts a non-negative integer n and a string s as parameters, and returns a string of s repeated exactly n times.
Examples (input -> output)

6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatString(5, "I"))
}
func RepeatString(n int, s string) string {
	return strings.Repeat(s, n)
}
