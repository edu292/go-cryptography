package main

import (
	"math/big"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	return new(big.Int).ModInverse(e, tot)
}
