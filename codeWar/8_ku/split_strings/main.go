package main

import "strings"

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
