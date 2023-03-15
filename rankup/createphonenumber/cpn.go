// Create Phone Number
// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.
// Example
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 uint
	fmt.Print("Введите номер по одной цифре: ")
	fmt.Fscan(os.Stdin, &n1)
	fmt.Fscan(os.Stdin, &n2)
	fmt.Fscan(os.Stdin, &n3)
	fmt.Fscan(os.Stdin, &n4)
	fmt.Fscan(os.Stdin, &n5)
	fmt.Fscan(os.Stdin, &n6)
	fmt.Fscan(os.Stdin, &n7)
	fmt.Fscan(os.Stdin, &n8)
	fmt.Fscan(os.Stdin, &n9)
	fmt.Fscan(os.Stdin, &n10)

	r := [10]uint{n1, n2, n3, n4, n5, n6, n7, n8, n9, n10}

	fmt.Println(CreatePhoneNumber1(r))
	fmt.Println(CreatePhoneNumber2(r))
}

func CreatePhoneNumber1(n [10]uint) string {
	result := ""
	i := 0
	for len(result) < 14 && i < 10 {
		if len(result) == 0 {
			result += "("
		} else if len(result) == 4 {
			result += ")"
		} else if len(result) == 5 {
			result += " "
		} else if len(result) == 9 {
			result += "-"
		} else {
			result += string(strconv.FormatUint(uint64(n[i]), 10))
			i += 1
		}
	}
	return result
}

func CreatePhoneNumber2(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}
