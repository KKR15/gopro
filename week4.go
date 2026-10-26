package main

import (
	"fmt"
	"math"
)

func standardDeviation(data []float64) float64 {
	n := float64(len(data))

	var sum float64
	for _, value := range data {
		sum += value
	}
	mean := sum / n

	var variance float64
	for _, value := range data {
		variance += math.Pow(value-mean, 2)
	}
	variance /= n

	return math.Sqrt(variance)
}

func main() {
	data := []float64{10, 20, 30, 40, 50}

	sd := standardDeviation(data)

	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Standard Deviation: %.2f\n", sd)
}
