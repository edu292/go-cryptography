package main

import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))
	randomKeyBytes := make([]byte, length)
	_, er := randReader.Read(randomKeyBytes)
	if er != nil {
		return "", er
	}

	randomKey := fmt.Sprintf("%x", randomKeyBytes)

	return randomKey, nil
}
