package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	const SECONDS_PER_DAY int = 86400
	const MARS_DISTANCE int = 62100000
	var company, trip string
	var speed, days, price int

	rand.Seed(time.Now().Unix())
	fmt.Printf("%s\n", strings.Repeat("=", 58))
	fmt.Printf("| %-16s | %-7s | %-4s | %-10s | %-5v |\n", "Spacelines", "Speed", "Days", "Trip", "Price")
	fmt.Printf("%s\n", strings.Repeat("=", 58))
	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed = rand.Intn(15) + 16
		days = MARS_DISTANCE / speed / SECONDS_PER_DAY
		price = speed + 20
		switch rand.Intn(2) {
		case 0:
			trip = "One-Way"
		case 1:
			trip = "Round-Trip"
			price *= 2
		}
		fmt.Printf("| %-16s | %v Km/h | %-4v | %-10s | $ %3v |\n", company, speed, days, trip, price)
	}
	fmt.Printf("%s\n", strings.Repeat("=", 58))
}
