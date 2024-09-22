package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  23,
		"second": 45,
	}

	floats := map[string]float64{
		"first":  39.78,
		"second": 13.26,
	}

	fmt.Printf("Non-generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
