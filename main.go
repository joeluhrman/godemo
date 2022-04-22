/*
 Joe Luhrman
 Dr. Binkley
 CS 451
 Pet demo 3/17/2022
 Mainly focussed on showing off goroutines
 Generates and calculates the average of multiple int arrays using Goroutines
*/

package main

import (
	"fmt"
	"math/rand"
)

const (
	N   = 10 // compiler will infer int
	LEN = 5
)

func main() { // curly brace must be on this line
	var arrays [N][LEN]int
	var avgs [N]float32

	// generating arrays
	for i := 0; i < N; i++ {
		data := make(chan [LEN]int)             // channel for int arrays
		go func() { data <- generateArray() }() // goroutine to generate each array concurrently
		arrays[i] = <-data
	}

	// calculating averages
	for i := 0; i < N; i++ {
		data := make(chan float32)                 // channel for floats
		go func() { data <- average(arrays[i]) }() // goroutine to calculate each average concurrently
		avgs[i] = <-data
	}

	// -----------------------------------------------------------------
	// OUTPUT
	fmt.Println("Arrays: ")
	for i := 0; i < N; i++ {
		fmt.Println("Array ", i, ": ", arrays[i])
	}

	fmt.Println("Averages: ")
	for i := 0; i < N; i++ {
		fmt.Println("Array ", i, ": ", avgs[i])
	}
}

// Generates an array of length LEN of ints from 0-10
func generateArray() [LEN]int {
	var arr [LEN]int

	for i := 0; i < LEN; i++ {
		arr[i] = rand.Intn(11) // generates ints from 0-10
	}

	return arr
}

// Calculates the average of an int array of length LEN
func average(arr [LEN]int) float32 {
	var average float32
	sum := 0

	for i := 0; i < LEN; i++ {
		sum += arr[i]
	}

	average = float32(sum) / float32(LEN)

	return average
}
