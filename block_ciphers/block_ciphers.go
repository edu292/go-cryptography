package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	var cipher cipher.Block
	var err error
	key := bytes.Repeat([]byte("-"), keyLen)
	switch cipherType {
	case typeAES:
		cipher, err = aes.NewCipher(key)
	case typeDES:
		cipher, err = des.NewCipher(key)
	default:
		return 0, errors.New("invalid cipher type")
	}
	if err != nil {
		return 0, err
	}

	return cipher.BlockSize(), nil
}
