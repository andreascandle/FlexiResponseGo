package utils

import (
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand
var once sync.Once

// initializeRand ensures randomness is seeded once.
func initializeRand() {
	once.Do(func() {
		seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
}

// GenerateTraceID creates a random trace ID.
func GenerateTraceID(length int) string {
	initializeRand()
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
