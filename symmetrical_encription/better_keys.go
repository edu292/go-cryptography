package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func keyToCipher(key string) (cipher.Block, error) {
	return aes.NewCipher([]byte(key))
}
