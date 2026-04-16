package main

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	return new(big.Int).Exp(c, d, n)
}
