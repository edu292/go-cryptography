package main

import (
	"errors"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i <= int(math.Exp2(24)); i++ {
		keyGuess := intToBytes(i)
		decryptTry := string(crypt(encrypted, keyGuess))
		if decryptTry == decrypted {
			return keyGuess, nil
		}
	}

	return nil, errors.New("key not found")
}
