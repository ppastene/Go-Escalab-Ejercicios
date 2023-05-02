package main

import "fmt"

func main() {
	var days int = 28
	var distance int = 56000000
	var totalHours int = days * 24

	var speed = distance / totalHours

	fmt.Println("La velocidad para viajar a Malacandra en 28 dias es:", speed)
}
