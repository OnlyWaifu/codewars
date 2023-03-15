// Convert a string to an array
// Write a function to split a string and convert it into an array of words.

// Examples (Input ==> Output):
// "Robin Singh" ==> ["Robin", "Singh"]

// "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(StringToArray("Пример словосочетания для проверки"))
}

func StringToArray(str string) []string {
	r := strings.Split(str, " ")
	return r
}
