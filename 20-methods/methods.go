package main

import "fmt"

type celsius float64
type kelvin float64
type farenheit float64

func (k kelvin) kelvinToCelsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) celsiusToFarenheit() farenheit {
	return farenheit((c * 9.0 / 5.0) + 32.0)
}

func main() {
	k := kelvin(300.0)
	c := celsius(30.0)
	fmt.Printf("Kelvin: %.1f°K\t to Celsius: %.1f°C\n", k, k.kelvinToCelsius())
	fmt.Printf("Celsius: %.1f°K\t to Farenheit: %.1f°C\n", c, c.celsiusToFarenheit())
	fmt.Printf("Kelvin: %.1f°K\t to Farenheit: %.1f°C\n", k, k.kelvinToCelsius().celsiusToFarenheit())
}
