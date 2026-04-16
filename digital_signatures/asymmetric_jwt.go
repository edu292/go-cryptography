package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	hasher := sha256.New()
	var hash [sha256.Size]byte

	hasher.Write([]byte(message))
	hasher.Sum(hash[:0])

	signatureBytes, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	jwt := message + "." + hex.EncodeToString(signatureBytes)
	return jwt, nil
}
