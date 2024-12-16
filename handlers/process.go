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

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format. Please verify input."})
		return
	}
	id := uuid.New().String()

	if err := dataStore.SaveReceipt(id, receipt); err != nil {
		log.Printf("Failed to save receipt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	points := CalculatePoints(receipt)
	log.Printf("Calculated Points:  %d", points)

	if err := dataStore.SavePoints(id, points); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save points"})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"id": id})
}
