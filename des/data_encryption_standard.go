package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"slices"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	b, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlaintest := padMsg(plaintext, des.BlockSize)
	ciphertext := make([]byte, des.BlockSize+len(paddedPlaintest))

	iv := ciphertext[:des.BlockSize]
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(b, iv)

	encrypter.CryptBlocks(ciphertext[des.BlockSize:], paddedPlaintest)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	plaintextLength := len(plaintext)
	fmt.Println(plaintextLength)
	nextMultiple := ((plaintextLength + blockSize - 1) / blockSize) * blockSize
	fmt.Println(nextMultiple)
	paddingSize := nextMultiple - plaintextLength

	return slices.Concat(plaintext, bytes.Repeat([]byte{0}, paddingSize))
}
