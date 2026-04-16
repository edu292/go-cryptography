package main

import (
	"math/big"
)

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, err := getBigPrime(keysize)
	if err != nil {
		return nil, nil
	}

	q, err := getBigPrime(keysize)
	if err != nil {
		return nil, nil
	}

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	n := big.NewInt(0)
	return n.Mul(p, q)
}
