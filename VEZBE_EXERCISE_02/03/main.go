/*-----------------------------------------------
Write a function with one variadic parameter that
finds the greatest number in a list of numbers.
------------------------------------------------*/

package main

import "fmt"

func max(brojevi ...int) int {
	var najveci int
	for _, naj := range brojevi {
		if naj > najveci {
			najveci = naj
		}
	}
	return najveci
}

func main() {
	najveci := max(1, 2, 3, 4, 5)
	fmt.Println(najveci)
}
