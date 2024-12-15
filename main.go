package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"
	"time"

	"github.com/dills122/receipt-rack/handlers"
	"github.com/dills122/receipt-rack/routes"
	"github.com/dills122/receipt-rack/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the port for the server from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback to default port if not set
	}

	serverUrl := fmt.Sprintf(":%s", port)

	// Initialize the data store (false for in-memory, true for Redis)
	dataStore := store.NewStore(false)

	// Initialize the handlers with the store
	handlers.Init(dataStore)

	// Create a Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Set up server with graceful shutdown
	server := &http.Server{
		Addr:    serverUrl,
		Handler: router,
	}

	// Run server in a goroutine to handle graceful shutdown
	go func() {
		log.Printf("Server is running on %s", serverUrl)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down gracefully...")
	// Timeout to wait for active connections to close
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server exited")
}
