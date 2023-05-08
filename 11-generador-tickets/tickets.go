package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	spacelines := [3]string{
		"Space Adventures\t",
		"SpaceX\t\t\t",
		"Virgin Galactic\t\t",
	}

	rand.Seed(time.Now().Unix())
	fmt.Println("Spaceline\t\t", "Days\t", "Trip Type\t", "Price\t")
	fmt.Println("----------------------------------------------------------")
	for i := 0; i < 10; i++ {
		// Choose a random number between 16 and 30, the value is meausred in Km/h
		speed := rand.Intn(30-16) + 16
		// With the distance in Km between the earth and mars we get the time in hours, then we divide it by 24 to get the days
		days := (62100000 / speed) / 24
		// 1 or 0, if 1 it will represent a Round-Trip ticket, else will be a One-Way
		roundTrip := rand.Intn(2)
		// With the value of speed we align it with the price by 20, so the range is between 36 and 50
		price := speed + 20
		// We print a random space travel company
		fmt.Printf(spacelines[rand.Intn(len(spacelines))])
		// We print the days
		fmt.Printf("%v\t", days)
		// If the trip is Round-Trip we multiply the final price by 2
		if roundTrip == 1 {
			fmt.Printf("%s\t$%v\n", "Round-Trip", price*2)
		} else {
			fmt.Printf("%s\t\t$%v\n", "One-Way", price)
		}
	}
	fmt.Println("----------------------------------------------------------")
}
