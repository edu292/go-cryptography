package main

import "math"

func alphabetSize(numBits int) float64 {
	return math.Exp2(float64(numBits))
}
