package main

import (
	"strings"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	cyphertext := make([]string, len(text))
	for i, v := range text {
		cyphertext[i] = getOffsetChar(v, key)
	}

	return strings.Join(cyphertext, "")
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	currentIndex := strings.Index(alphabet, string(c))
	if currentIndex == -1 {
		return ""
	}

	newIndex := mod(currentIndex+offset, len(alphabet))

	return string(alphabet[newIndex])
}
