package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}

	plaintext, err = gcm.Open(plaintext, nonce, ciphertext, []byte{})
	if err != nil {
		return nil, err
	}
	return
}
