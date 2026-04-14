package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	bytes := make([]string, len(b))
	for i, v := range b {
		bytes[i] = fmt.Sprintf("%02x", v)
	}

	return strings.Join(bytes, ":")
}

func getBinaryString(b []byte) string {
	bytes := make([]string, len(b))
	for i, v := range b {
		bytes[i] = fmt.Sprintf("%08b", v)
	}

	return strings.Join(bytes, ":")
}
