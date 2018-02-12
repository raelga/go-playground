package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(x int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(x)
}

func main() {
	fmt.Printf("Here you have a random number %d\n", random(100))
}
