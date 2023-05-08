package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int
	var counter int = 0
	fmt.Print("Type a number between 1 and 100: ")
	fmt.Scanln(&number)
	rand.Seed(time.Now().UnixNano())
	if number < 1 || number > 100 {
		panic("I TOLD YOU A NUMBER BETWEEN 1 AND 100!")
	}
	for {
		var random int = rand.Intn(100) + 1
		if number == random {
			switch counter {
			case 1:
				fmt.Printf("The code took %d try to match the number", counter)
			default:
				fmt.Printf("The code took %d tries to match the number", counter)
			}
			break
		}
		counter++
	}
}
