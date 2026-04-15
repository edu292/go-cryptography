package main

import (
	"fmt"
	"slices"
)

func feistel(msg []byte, roundKeys [][]byte) []byte {
	sideSize := len(msg) / 2
	lhs := msg[:sideSize]
	rhs := msg[sideSize:]

	fmt.Printf("%d", len(roundKeys))
	for _, key := range roundKeys {
		nextLhs := rhs
		rhs = xor(hash(rhs, key, sideSize), lhs)

		lhs = nextLhs
	}

	return slices.Concat(rhs, lhs)
}
