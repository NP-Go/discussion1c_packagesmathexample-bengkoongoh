package main

import (
	"fmt"
	"math"
)

func main() {
	q := 34.0
	r := 8.0

	s := math.Mod(q, r)
	fmt.Println(s)

	// modWithFloat := q % 2
	// fmt.Println(modWithFloat)
	// operator % not defined on q (variable of type float64)

	// convert to integer
	modWithInteger := int(q) % 2
	fmt.Println(modWithInteger)
}
