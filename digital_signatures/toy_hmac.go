package main

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
)

func hmac(message, key string) string {
	hasher := sha256.New()
	middleIndex := len(key) / 2
	keyFirstHalf := []byte(key[:middleIndex])
	keySecondHalf := key[middleIndex:]

	var hash [sha256.Size]byte

	hasher.Write([]byte(keySecondHalf + message))
	hasher.Sum(hash[:0])
	hasher.Reset()
	hasher.Write(slices.Concat(keyFirstHalf, hash[:]))

	hasher.Sum(hash[:0])
	return hex.EncodeToString(hash[:])
}
