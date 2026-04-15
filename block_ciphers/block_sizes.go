package main

import "bytes"

func padWithZeros(block []byte, desiredSize int) []byte {
	missingLength := desiredSize - len(block)
	block = append(block, bytes.Repeat([]byte{0}, missingLength)...)
	return block
}
