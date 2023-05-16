// A closure is a function that references variables outside of its scope
// The closure is a nested anonymous function,
// and can hold the values of the variables even after the scope is destroyed
package main

import "fmt"

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(a, b int) func() int {
	return func() int {
		return a + b
	}
}

func main() {
	// It helps with the data insulation
	num1 := incrementor()
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())
	num2 := incrementor()
	fmt.Println(num2())
	fmt.Println(num2())
	// We declare variables a and b, and sum which is assigned the sum() function
	a := 3
	b := 5
	// Calling the outter function
	result := sum(a, b)
	// Calling the inner function
	fmt.Println(result())
	a++
	b++
	// So, despite of incrementing a and b variables, why result() didnt change its result value?
	// Because sum() holds a copy of a and b, not the references
	fmt.Println(result())
}
