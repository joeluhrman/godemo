// Joe Luhrman
// Dr. Binkley
// CS 451
// Pet demo draft 3/17/2022
// Mainly focussed on showing off goroutines

package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	return a + b
}

func display(a int) {
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

func main() {
	a := 5
	b := 7
	c := 10
	d := 9

	go fmt.Println(add(a, b))

	display(add(c, d))

}
