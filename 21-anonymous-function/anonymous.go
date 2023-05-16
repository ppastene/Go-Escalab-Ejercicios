package main

import "fmt"

// An anonymous function can be assigned in a variable and be called
var f = func() {
	fmt.Println("Go is an amazing language!")
}

func main() {
	f()
	// Also can be declarated in the function's scope
	f := func(text string) {
		fmt.Println(text)
	}
	f("It can be used to create amazing projects")
	// And can be created and be invoked at the same time
	func() {
		fmt.Println("So give it a try, you wont regret it")
	}()
}
