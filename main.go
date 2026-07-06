package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Pi = 22.0 / 7

func luas(
	r float64,
) float64 {
	return Pi * r * r
}

func keliling(
	r float64,
) float64 {
	return Pi * r * 2
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

	fmt.Printf("lingkaran dengan r = %v\n", r)
	fmt.Printf("  luas    : %v\n", luas(r))
	fmt.Printf("  keliling: %v\n", keliling(r))
}
