// Beginner - Lost Without a Map
// Given an array of integers, return a new array with each value doubled.

// For example:

// [1, 2, 3] --> [2, 4, 6]

package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2, n3 int
	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &n1)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &n2)

	fmt.Print("Введите третье число: ")
	fmt.Fscan(os.Stdin, &n3)

	r := []int{n1, n2, n3}

	fmt.Println(Maps(r))
}

func Maps(x []int) []int {
	r := make([]int, len(x))
	for i := range x {
		r[i] = x[i] * 2
	}
	return r
}
