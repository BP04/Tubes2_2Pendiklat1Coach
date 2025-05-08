package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"search"
)

func main() {
	// Load elements.json
	dataPath, err := filepath.Abs("data/elements.json")
	if err != nil {
		log.Fatal("Error finding elements.json:", err)
	}
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatal("Error reading elements.json:", err)
	}
	var elements []search.Element
	if err := json.Unmarshal(data, &elements); err != nil {
		log.Fatal("Error parsing elements.json:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// API endpoint for search
	r.POST("/api/search", func(c *gin.Context) {
		var request struct {
			Element    string `json:"element"`
			Algorithm  string `json:"algorithm"`
			Mode       string `json:"mode"`
			MaxRecipes int    `json:"maxRecipes"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Perform search
		recipes, time, nodesVisited, err := search.FindRecipes(
			elements,
			request.Element,
			request.Algorithm,
			request.Mode,
			request.MaxRecipes,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond
		c.JSON(http.StatusOK, gin.H{
			"recipes":      recipes,
			"time":         time,
			"nodesVisited": nodesVisited,
		})
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}