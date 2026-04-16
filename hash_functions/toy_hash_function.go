package main

import "math/bits"

func hash(input []byte) [4]byte {
	digest := [4]byte{}
	for i, b := range input {
		b = bits.RotateLeft8(b, 3)
		b = b << 2
		digest[i%4] ^= b
	}

	return digest
}
