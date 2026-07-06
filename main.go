package main

import (
	"fmt"
	"math"
)

func luas(
	r float32,
) float32 {
	return math.Pi * r * r
}

func keliling(
	r float32,
) float32 {
	return math.Pi * r * 2
}

func main() {
	fmt.Printf("lingkaran dengan r = 7 \n")
	fmt.Printf("  luas    : %f\n", luas(7))
	fmt.Printf("  keliling: %f\n", keliling(7))
}
