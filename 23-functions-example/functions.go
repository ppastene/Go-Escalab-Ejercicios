package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

// Both fakeSensor() and realSensor() will return a value of type kelvin
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 //TODO: Implement real sensor
}

// This function accetps an integer and a function as parameters
// The sensor name will hold the return value of the assigned function
// Also the function's parameters can be written like this:
// type sensor func() kelvin
// func measureTemperature(samples int, s sensor) {}
func measureTemperature(samples int, s func() kelvin) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v grados Kelvin\n", k)
		time.Sleep(time.Second)
	}
}

// This is a closure function, which uses an anonymous function as a return
// and the inner function holds the arguments from the outter function
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	// Both sensor1 and sensor2 variables have a function as a value,
	// so it will have their return values
	sensor1 := fakeSensor
	fmt.Println(sensor1())

	sensor2 := realSensor
	fmt.Println(sensor2())

	// The fakeSensor function passes as an argument to measureTemperature
	// So the return value of fakeSensor will be used inside the function
	measureTemperature(3, fakeSensor)

	// Despite of changing the offset value, the calibrate function will return the same value
	// That's because the function calibrate() is managing the arguments as a value (a copy) instead as a reference
	// So it doesnt matter if the value of offset changes, the calibrate() have a copy of the value of the variable
	var offset kelvin = 5
	sensor3 := calibrate(realSensor, offset)
	fmt.Println(sensor3())
	offset++
	fmt.Println(sensor3())
}
