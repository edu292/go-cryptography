package main

import "math"

// nonceStrength returns the number of bits of entropy in the nonce.
func nonceStrength(nonce []byte) int {
	return int(math.Exp2(float64(len(nonce) * 8)))
}
