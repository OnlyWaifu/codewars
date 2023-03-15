//			Square(n) Sum
// Complete the square sum function so that it squares each number passed into it and then sums the results together.
// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9

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

	fmt.Println(SquareSum(r))
}

func SquareSum(numbers []int) int {
	r := 0
	for i, a := range numbers {
		numbers[i] = a * a
		r += numbers[i]
	}
	return r
}
