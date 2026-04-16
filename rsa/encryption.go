package main

import (
	"math/big"
)

func encrypt(m, e, n *big.Int) *big.Int {
	return new(big.Int).Exp(m, e, n)
}
