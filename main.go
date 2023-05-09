package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	})

	type Item struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Image    string `json:"image"`
	}

	// Add some initial items to the vending machine
	items := []Item{
		{ID: 1, Name: "Mirinda", Price: 50, Quantity: 5, Image: "data:image/jpeg;base64,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"},
		{ID: 2, Name: "Darkfantasy", Price: 100, Quantity: 5, Image: "data:image/jpeg;base64,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"},
		{ID: 3, Name: "Lays", Price: 20, Quantity: 5, Image: "data:image/jpeg;base64,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"},
		{ID: 4, Name: "Littleheart", Price: 30, Quantity: 5, Image: "data:image/jpeg;base64,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"},
	}

	// Get all items from the vending machine
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)

	})

	// Start the server
	router.Run(":8081")
}