package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func checksumMatches(message string, checksum string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(message))

	var hash [sha256.Size]byte
	hasher.Sum(hash[:0])

	return hex.EncodeToString(hash[:]) == checksum
}
