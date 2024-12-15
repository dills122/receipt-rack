package routes

import (
	"github.com/dills122/receipt-rack/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the API.
func RegisterRoutes(router *gin.Engine) {
	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetReceiptPoints)
}
