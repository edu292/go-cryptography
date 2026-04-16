package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	timeCost    uint32 = 3
	memoryCost  uint32 = 32 * 1024
	parallelism uint8  = 4
	keyLen             = 32
	saltLen            = 16
)

func hashPassword(password string) (string, error) {
	var salt [saltLen]byte
	if _, err := rand.Read(salt[:]); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt[:], timeCost, memoryCost, parallelism, keyLen)
	b64 := base64.RawStdEncoding

	phc := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memoryCost, timeCost, parallelism,
		b64.EncodeToString(salt[:]), b64.EncodeToString(hash))

	return phc, nil
}

func checkPasswordHash(password, hash string) bool {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" || parts[2] != "v=19" {
		return false
	}

	parameters := strings.Split(parts[3], ",")
	if len(parameters) != 3 {
		return false
	}

	var m, t uint32
	var p uint8
	for _, parameter := range parameters {
		keyValue := strings.Split(parameter, "=")
		if len(keyValue) != 2 {
			return false
		}

		key := keyValue[0]
		value64, err := strconv.ParseUint(keyValue[1], 10, 32)
		if err != nil {
			return false
		}

		value := uint32(value64)

		switch key {
		case "m":
			m = value
		case "t":
			t = value
		case "p":
			if value > 255 {
				fmt.Println("e")
				return false
			}

			p = uint8(value)
		default:
			return false
		}
	}

	if p == 0 || t == 0 || m == 0 {
		return false
	}

	b64 := base64.RawStdEncoding
	salt, err := b64.DecodeString(parts[4])
	if err != nil {
		return false
	}
	want, err := b64.DecodeString(parts[5])
	if err != nil {
		return false
	}

	got := argon2.IDKey([]byte(password), salt, t, m, p, uint32(len(want)))
	return subtle.ConstantTimeCompare(want, got) == 1
}
