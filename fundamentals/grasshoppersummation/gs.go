// Grasshopper - Summation
// Summation
// Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0.

// For example (Input -> Output):

// 2 -> 3 (1 + 2)
// 8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)

package main

import (
	"fmt"
	"os"
)

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &n)

	fmt.Println(Summation1(n))
	fmt.Println(Summation2(n))
}

func Summation1(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		r += i
	}
	return r
}

func Summation2(n int) int {
	return n * (n + 1) / 2
}
