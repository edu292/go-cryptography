package main

var englishFreq = map[rune]int{
	'e': 12,
	't': 9,
	'a': 8,
	'o': 7,
	'i': 6,
	'n': 6,
	's': 6,
	'h': 6,
	'r': 5,
	'd': 4,
	'l': 4,
	'c': 2,
	'u': 2,
	'm': 2,
	'w': 2,
	'f': 2,
	'g': 2,
	'y': 1,
	'p': 1,
	'b': 1,
	'v': 0,
	'k': 0,
	'j': 0,
	'x': 0,
	'q': 0,
	'z': 0,
}

func crack(ciphertext string) string {
	bestShift := 0
	maxScore := 0

	for shift := 0; shift < 26; shift++ {
		score := 0

		for _, char := range ciphertext {
			shifted := char - rune(shift)
			if shifted < 'a' {
				shifted += 26
			}
			score += englishFreq[shifted]
		}

		if score > maxScore {
			maxScore = score
			bestShift = shift
		}
	}

	return decrypt(ciphertext, bestShift)
}
