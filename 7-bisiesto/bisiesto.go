package main

import "fmt"

func main() {
	var year int = 1990
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Println("Año bisiesto")
	} else {
		fmt.Println("Año no bisiesto")
	}
}
