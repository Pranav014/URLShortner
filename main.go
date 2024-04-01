package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"

	"github.com/gin-gonic/gin"
)

type longUrlRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shorturl"`
	LongURL  string `json:"longurl"`
}

var URLMap = make(map[string]string)

var predefinedString = "my_salt_string"

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func shortenURL(c *gin.Context) {
	var req longUrlRequest

	// Bind the JSON data from the request body to the ShortenRequest struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error parsing the JSON data, return an error response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create a UUID for the long URL
	//newUUID := uuid.New()

	trimmedHash := createHash(req.LongURL)
	appendedString := req.LongURL
	// Add the UUID to the URLMap
	// Check if the key exists in the map
	for {
		if value, ok := getFromRedis(trimmedHash); ok == "" {
			fmt.Println("value of hash", value)
			fmt.Println("value of trimmed hash", trimmedHash)
			// Key exists, get the value
			appendedString = appendedString + predefinedString
			trimmedHash = createHash(appendedString)
		} else {
			// Key doesn't exist, break out of the loop
			storeInRedis(trimmedHash, req.LongURL)
			break
		}
	}

	resp := ShortenResponse{
		ShortURL: trimmedHash,
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	if longURL, ok := getFromRedis(shortURL); ok == "" {
		c.Redirect(http.StatusMovedPermanently, longURL)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
	}
}

// Function to store trimmedhash and value in redis
func storeInRedis(trimmedHash string, longURL string) {
	err := client.Set(ctx, trimmedHash, longURL, 0).Err()
	if err != nil {
		panic(err)
	}
}

// Function to get value from redis
func getFromRedis(trimmedHash string) (string, string) {
	val, err := client.Get(ctx, trimmedHash).Result()
	if err != nil {
		return "", err.Error()
	}
	return val, ""

}

func createHash(longURL string) string {
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

//create hash function with salt
func createHashWithSalt(longURL string, salt string) string {
	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write UUID string to the hash
	hash.Write([]byte(longURL + salt))

	// Get the hashed bytes
	hashedBytes := hash.Sum(nil)

	// Convert hashed bytes to a hex string
	hashedStr := hex.EncodeToString(hashedBytes)
	fmt.Println("Inside create hash with salt")

	trimmedHash := hashedStr[:8]

	return trimmedHash

}

func main() {
	router := gin.Default()
	router.POST("/shortenURL", shortenURL)
	router.GET("/:shortURL", redirectURL)

	router.Run("localhost:8080")

}
