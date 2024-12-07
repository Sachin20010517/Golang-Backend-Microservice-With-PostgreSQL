package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateUniqueID(prefix string) string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%s-%x", prefix, b)
}
