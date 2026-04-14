package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	bytes := strings.Join(strings.Split(s, ":"), "")
	return hex.DecodeString(bytes)
}
