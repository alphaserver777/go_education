/*Create a function that accepts a parameter representing a name and returns the message: "Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]
*/

package main

import "fmt"

func main() {
	Greet("Slava")
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
