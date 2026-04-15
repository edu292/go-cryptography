package main

import "errors"

var lookup = [][]byte{
	{0b00, 0b10, 0b01, 0b11},
	{0b10, 0b00, 0b11, 0b01},
	{0b01, 0b11, 0b00, 0b10},
	{0b11, 0b01, 0b10, 0b00},
}

func sBox(b byte) (byte, error) {
	if b > 0b1111 {
		return 0, errors.New("invalid input")
	}

	row := b >> 2
	col := b & 0b0011

	return lookup[row][col], nil
}
