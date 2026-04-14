package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	index := int(bits)
	if index >= len(base8Alphabet) {
		return ""
	}

	return string(base8Alphabet[index])
}
