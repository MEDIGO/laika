package store

import "time"

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
