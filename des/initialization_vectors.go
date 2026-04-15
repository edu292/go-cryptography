package main

import "crypto/rand"

func generateIV(length int) ([]byte, error) {
	iv := make([]byte, length)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}

	return iv, nil
}
