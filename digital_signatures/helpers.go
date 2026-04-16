package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

func verifyECDSAMessage(token string, publicKey *ecdsa.PublicKey) error {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return errors.New("invalid token sections")
	}
	sig, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(parts[0]))

	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}

