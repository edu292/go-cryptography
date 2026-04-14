package main

func crypt(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, v := range plaintext {
		ciphertext[i] = v ^ key[i]
	}

	return ciphertext
}
