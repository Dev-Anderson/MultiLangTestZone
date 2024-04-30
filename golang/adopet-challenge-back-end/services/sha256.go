package services

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encoder(s string) string {
	str := sha256.Sum224([]byte(s))
	return fmt.Sprintf("%x", str)
}
