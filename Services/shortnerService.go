package Services

import (
	"awesomeProject/helper"
	"fmt"
)

const predefinedString = "my_salt_string"

// UrlShortener Function to shorten the URL
func UrlShortener(longURL string) (string, error) {
	// This function will be used to shorten the URL
	trimmedHash := helper.CreateHash(longURL)
	appendedString := longURL
	// Add the UUID to the URLMap
	// Check if the key exists in the map
	for {
		if value, ok := helper.GetFromRedis(trimmedHash); ok == "" {
			fmt.Println("value of hash", value)
			fmt.Println("value of trimmed hash", trimmedHash)
			// Key exists, get the value
			appendedString = appendedString + predefinedString
			trimmedHash = helper.CreateHash(appendedString)
		} else {
			// Key doesn't exist, break out of the loop
			helper.StoreInRedis(trimmedHash, longURL)
			return trimmedHash, nil
		}

	}
}

func RedirectURL(shortURL string) (string, error) {
	longURL, ok := helper.GetFromRedis(shortURL)
	if ok == "" {
		return longURL, nil
	}
	return "", fmt.Errorf("URL not found")
}

func deleteURL(shortURL string) {
	// This function will be used to delete the URL
	helper.DeleteFromRedis(shortURL)
}

func DeleteURL(url string) error {
	// Delete the URL from Redis
	err := helper.DeleteFromRedis(url)
	if err != nil {
		// If there was an error deleting the URL from Redis, return an error response
		return fmt.Errorf("failed to delete URL: %w", err)
	}

	// If deletion was successful, return a success message
	return nil
}
