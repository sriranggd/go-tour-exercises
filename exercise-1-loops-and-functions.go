package main

import (
	"fmt"
	"math"
)

func Sqrtv1(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (z * 2)
		fmt.Println(z)
	}

	return z
}

func Sqrtv2(x float64) float64 {
	z := 1.0
	difference := 0.1
	for difference > 0.00001 {
		var newz = z - ((z*z - x) / (z * 2))
		difference = math.Abs(newz - z)
		z = newz
		fmt.Println(z)
	}

	return z
}
