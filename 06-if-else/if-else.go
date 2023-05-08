package main

import "fmt"

func main() {
	var action = "Go to the left"

	if action == "Go to the left" {
		fmt.Println("You arrived to an old house")
	} else if action == "Go to the right" {
		fmt.Println("You arrived to a river")
	} else {
		fmt.Println("You don't know where you are")
	}
}
