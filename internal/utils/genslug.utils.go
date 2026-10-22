package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	slug := make([]byte, length)
	for i := range slug {
		slug[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(slug)
}

func ValidateSlug(slug string) (bool, error) {
	if len(slug) < 3 || len(slug) > 50 {
		return false, ErrLengthSlug
	}
	reserveWord := map[string]bool{
		"login":     true,
		"register":  true,
		"api":       true,
		"dashboard": true,
	}
	if reserveWord[strings.ToLower(slug)] {
		return false, ErrReservedWord
	}
	validatingWithRegex := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !validatingWithRegex.MatchString(slug) {
		return false, ErrFormatSlug
	}
	return true, nil
}
