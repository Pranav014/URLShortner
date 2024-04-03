package main

import (
	"awesomeProject/Services"
	"github.com/gin-gonic/gin"

	"net/http"
)

type longUrlRequest struct {
	LongURL string `json:"long_url" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"shorturl"`
	LongURL  string `json:"longurl"`
}

func shortenURL(c *gin.Context) {
	var req longUrlRequest

	// Bind the JSON data from the request body to the ShortenRequest struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error parsing the JSON data, return an error response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Call the URLShortener function from the Services package
	trimmedHash, err := Services.UrlShortener(req.LongURL)
	if err != nil {
		// If there's an error while shortening the URL, return an internal server error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := ShortenResponse{
		ShortURL: trimmedHash,
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func redirectURLendpoint(c *gin.Context) {
	shortURL := c.Param("shortURL")
	longURL, err := Services.RedirectURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, longURL)
	}
}

// Endpoint to Delete
func deleteURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	err := Services.DeleteURL(shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully"})

}

func main() {
	router := gin.Default()
	router.POST("/shortenURL", shortenURL)
	router.GET("/:shortURL", redirectURLendpoint)
	router.DELETE("/:shortURL", deleteURL)
	router.Run("localhost:8080")

}
