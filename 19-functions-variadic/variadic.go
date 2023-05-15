// A variadic function is just a way to declare a function that accepts an unlimited number of argeuments
// An example could be fmt.Println() which accepts n ammount of arguments to print on console

package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(5, 10, 35))
	fmt.Println(sum(9, 7, 12, 6, 42))
	nums := []int{15, 50, 25, 5, 5}
	fmt.Println(sum(nums...))
}
