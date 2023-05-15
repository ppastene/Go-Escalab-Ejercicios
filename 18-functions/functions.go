package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFarenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFarenheit(k float64) float64 {
	return kelvinToCelsius(celsiusToFarenheit(k))
}

func main() {
	kelvin := 300.0
	celsius := 30.0
	fmt.Printf("Kelvin: %.1f°K\t to Celsius: %.1f°C\n", kelvin, kelvinToCelsius(kelvin))
	fmt.Printf("Celsius: %.1f°K\t to Farenheit: %.1f°C\n", celsius, celsiusToFarenheit(celsius))
	fmt.Printf("Kelvin: %.1f°K\t to Farenheit: %.1f°C\n", kelvin, kelvinToFarenheit(kelvin))
}
