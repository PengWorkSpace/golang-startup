package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var numArrys [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		numArrys[i] = rand.Intn(100)
	}

	fmt.Println("交换前数组", numArrys)

	arraylen := len(numArrys)
	var temp int
	for i := 0; i < arraylen/2; i++ {
		temp = numArrys[arraylen-1-i]
		numArrys[arraylen-1-i] = numArrys[i]
		numArrys[i] = temp
	}
	fmt.Println("交换后数组", numArrys)
}
