package routes

import (
	"github.com/dills122/receipt-rack/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetReceiptPoints)
}
