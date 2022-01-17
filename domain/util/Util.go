package util

import (
	"fmt"
	"math/rand"
	"time"
)

func ConvertTimeToString(value *time.Time) string {
	return value.Format(time.RFC3339)
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
