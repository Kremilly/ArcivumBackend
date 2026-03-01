package generator

import (
	"crypto/rand"
	"encoding/base64"
)

func Token(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}