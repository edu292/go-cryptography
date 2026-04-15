package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

// keysArePaired verifies if the public and private keys are paired using ECDSA.
func keysArePaired(pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) bool {
	msg := "a test message"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return false
	}

	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}

