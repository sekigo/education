package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, el := range m {
		s += el
	}
	return s

}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, el := range m {
		s += el
	}
	return s

}

func SumIntAndFloat[K comparable, V AllowedTypes](m map[K]V) V {
	var s V
	for _, el := range m {
		s += el
	}
	return s
}

type AllowedTypes interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntAndFloat[string, int64](ints),
		SumIntAndFloat[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntAndFloat(ints),
		SumIntAndFloat(floats))

}
