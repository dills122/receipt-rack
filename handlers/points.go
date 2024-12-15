package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	points, exists := dataStore.GetPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No points found for that ID."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points.Points})
}
