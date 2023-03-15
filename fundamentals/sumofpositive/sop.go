// Sum of positive
// You get an array of numbers, return the sum of all of the positives ones.

// Example [1,-4,7,12] => 1 + 7 + 12 = 20

// Note: if there is nothing to sum, the sum is default to 0.

package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2, n3, n4 int
	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &n1)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &n2)

	fmt.Print("Введите третье число: ")
	fmt.Fscan(os.Stdin, &n3)

	fmt.Print("Введите четвёртое число: ")
	fmt.Fscan(os.Stdin, &n4)

	r := []int{n1, n2, n3, n4}

	fmt.Println(PositiveSum(r))
}

func PositiveSum(numbers []int) int {
	r := 0
	for i, a := range numbers {
		if numbers[i] < 0 {
			a = 0
		}
		r += a
	}
	return r
}
