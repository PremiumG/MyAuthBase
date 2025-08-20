package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

// var MagicLinks = sync.Map{} // email -> token

type MagicToken struct {
	UserEmail string
	Expiry    time.Time
}

// Global storage (use Redis/DB in production)
var MagicTokens sync.Map

func CreateMagicLink(userEmail string) (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	// Store with expiration (use Redis/DB in production)
	expiry := time.Now().Add(5 * time.Minute)
	MagicTokens.Store(token, MagicToken{
		UserEmail: userEmail,
		Expiry:    expiry,
	})

	baseURL := "localhost:8080"
	// Only token in URL, no email
	return fmt.Sprintf("%s/verifymagicregister?token=%s", baseURL, token), nil

}

func init() {
	StartTokenCleanup()
}

func StartTokenCleanup() {
	ticker := time.NewTicker(5 * time.Minute) // Clean every 5 minutes
	go func() {
		for range ticker.C {
			cleanupExpiredTokens()
		}
	}()
}

func cleanupExpiredTokens() {
	now := time.Now()
	MagicTokens.Range(func(key, value interface{}) bool {
		token := value.(MagicToken)
		if now.After(token.Expiry) {
			MagicTokens.Delete(key)
		}
		return true
	})
}
