package main

import "fmt"

func main() {
	var counter int = 0
	fmt.Println("This is a loop, equal to the while loop in other languages")
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("This is another loop, similar to the for loop with its own counter")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("You can declare loops without a conditional, but since is not recommended we can put inside a conditional to break the loop")
	counter = 0
	for {
		fmt.Println(counter)
		counter++
		if counter > 10 {
			fmt.Println("We break here to not make an infinite loop")
			break
		}
	}
}
