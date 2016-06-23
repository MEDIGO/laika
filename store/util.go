package store

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func Bool(b bool) *bool {
	x := b
	return &x
}

func Int(i int64) *int64 {
	x := i
	return &x
}

func String(s string) *string {
	x := s
	return &x
}

func Time(t time.Time) *time.Time {
	x := t
	return &x
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
