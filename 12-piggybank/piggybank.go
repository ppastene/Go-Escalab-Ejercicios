package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cash []float64
	var piggybank float64 = 0.0
	cash = append(cash, 0.05)
	cash = append(cash, 0.10)
	cash = append(cash, 0.25)
	for piggybank < 20.0 {
		rand.Seed(time.Now().Unix())
		piggybank += cash[rand.Intn(len(cash))]
		fmt.Printf("There's %.2f of cash in the piggybank\n", piggybank)
	}
}
