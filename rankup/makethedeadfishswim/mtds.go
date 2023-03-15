// Make the Deadfish Swim
// Write a simple parser that will parse and run Deadfish.

// Deadfish has 4 commands, each 1 character long:

// i increments the value (initially 0)
// d decrements the value
// s squares the value
// o outputs the value into the return array
// Invalid characters should be ignored.

// Parse("iiisdoso") == []int{8, 64}

package main

import (
	"fmt"
)

func main() {
	var n string

	fmt.Print("Введите набор комманд (i, d, s, o): ")
	fmt.Scan(&n)

	fmt.Println(Parse(n))
}

func Parse(data string) []int {
	c := []int{}
	r := 0
	for _, m := range data {
		switch m {
		case 'i':
			r += 1
		case 'd':
			r -= 1
		case 's':
			r *= r
		case 'o':
			c = append(c, r)
		}
	}
	return c
}
