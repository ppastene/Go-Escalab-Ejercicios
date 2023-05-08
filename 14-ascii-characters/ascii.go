package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var exclamation rune = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, exclamation)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, exclamation)

	var message string = "shalom"
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c %v\n", message[i], message[i])
	}
}
