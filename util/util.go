package util

import (
	"crypto/rand"
	"encoding/hex"
)

func Salt() string {
	return Randstr(16)
}

func Token() string {
	return Randstr(13)
}

func Randstr(len int) string {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:len]
}
