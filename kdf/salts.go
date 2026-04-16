package main

import (
	"crypto/rand"
	"crypto/sha256"
	"slices"
)

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, 0, length)
	rand.Read(salt)

	return salt, nil
}

func hashPassword(password, salt []byte) []byte {
	hasher := sha256.New()
	hasher.Write(slices.Concat(password, salt))
	var hash [sha256.Size]byte
	hasher.Sum(hash[:0])

	return hash[:]
}
