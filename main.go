package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var amount int
	var signs = [2]string{"-", "+"}
	for {
		fmt.Print("Amount of random steps: ")
		fmt.Scanln(&amount)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < amount; i++ {
			fsign := rand.Intn(2)
			first := rand.Intn(24) + 1
			ssign := rand.Intn(2)
			second := rand.Intn(24) + 1
			fmt.Print(signs[fsign], first, ", ", signs[ssign], second, "\n")
		}
	}
}
