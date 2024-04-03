package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CreateHash(longURL string) string {
	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write UUID string to the hash
	hash.Write([]byte(longURL))

	// Get the hashed bytes
	hashedBytes := hash.Sum(nil)

	// Convert hashed bytes to a hex string
	hashedStr := hex.EncodeToString(hashedBytes)
	fmt.Println("Inside create hash")
	trimmedHash := hashedStr[:8]

	return trimmedHash
}
