package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

// generateRandomString cria uma string longa com um conjunto limitado de caracteres para o benchmark.
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

var longString = generateRandomString(10000)

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring(longString)
	}
}