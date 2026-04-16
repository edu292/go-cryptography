package main

import (
	"crypto/rand"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	one := big.NewInt(1)
	p1 := new(big.Int).Sub(p, one)
	q1 := new(big.Int).Sub(q, one)

	return new(big.Int).Mul(p1, q1)
}

func getE(tot *big.Int) *big.Int {
	one := big.NewInt(1)
	four := big.NewInt(4)
	maxRand := new(big.Int).Sub(tot, big.NewInt(5))

	for {
		e, err := rand.Int(randReader, maxRand)
		if err != nil {
			return nil
		}

		e.Add(e, four)
		if gcd(e, tot).Cmp(one) == 0 {
			return e
		}
	}
}
