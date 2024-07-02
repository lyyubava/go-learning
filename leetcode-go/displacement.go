package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var fn func(float64) float64
	var a, v0, s0, s, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)
	for true {
		fmt.Print("Enter initial time: ")
		fmt.Scan(&t)
		fn = GenDisplaceFn(a, v0, s0)
		s = fn(t)
		fmt.Printf("%.4f\n", s)
	}
}
