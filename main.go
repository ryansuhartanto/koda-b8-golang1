package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func luas(
	r float64,
) float64 {
	return math.Pi * r * r
}

func keliling(
	r float64,
) float64 {
	return math.Pi * r * 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("silahkan input jari-jari (r) lingkaran: ")
	scanner.Scan()
	r, err := strconv.ParseFloat(scanner.Text(), 32)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Kesalahan input!\n")
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	fmt.Printf("lingkaran dengan r = 7 \n")
	fmt.Printf("  luas    : %f\n", luas(r))
	fmt.Printf("  keliling: %f\n", keliling(r))
}
