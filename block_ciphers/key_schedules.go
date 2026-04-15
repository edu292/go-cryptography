package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	b := byte(roundNumber)
	derivedKey := [len(masterKey)]byte{}
	for i, v := range masterKey {
		derivedKey[i] = v ^ b
	}

	return derivedKey
}
