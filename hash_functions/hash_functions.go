package main

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{hash: sha256.New()}
}

func (h *hasher) Write(data string) (int, error) {
	return h.hash.Write([]byte(data))
}

func (h *hasher) GetHex() string {
	return hex.EncodeToString(h.hash.Sum([]byte{}))
}
