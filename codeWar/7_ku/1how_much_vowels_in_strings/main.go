/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/

package main

import "strings"

func main() {
	GetCount("Vladivostok")

}

func GetCount(str string) int {
	vowels := "aeiou"
	count := 0

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func GetCount2(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
