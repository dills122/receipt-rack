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
	"github.com/dills122/receipt-rack/middleware"
	"github.com/dills122/receipt-rack/routes"
	"github.com/dills122/receipt-rack/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const Default_Port = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = Default_Port
	}
	serverUrl := fmt.Sprintf(":%s", port)

	dataStore := store.NewStore(false)
	handlers.Init(dataStore)

	router := gin.New()
	router.Use(middleware.SecurityHeaders())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	routes.RegisterRoutes(router)

	server := &http.Server{
		Addr:    serverUrl,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server failed: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	fmt.Println("Received shutdown signal, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown Failed:%+v\n", err)
	}

	fmt.Println("Server exited gracefully")
}
