package handlers

import (
	"log"
	"net/http"

	"github.com/dills122/receipt-rack/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Bind JSON to the Receipt model
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format. Please verify input."})
		return
	}

	// Generate unique ID for the receipt
	id := uuid.New().String()

	// Save the receipt in the store
	if err := dataStore.SaveReceipt(id, receipt); err != nil {
		log.Printf("Failed to save receipt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Calculate points (example: points based on items)
	points := len(receipt.Items) * 10

	// Save points for the receipt
	if err := dataStore.SavePoints(id, points); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "points": points})
}
