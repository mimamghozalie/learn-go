package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// HelloResponse struct
type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(c *gin.Context) {
	response := HelloResponse{
		Message: "Hello, World!",
	}

	c.JSON(200, response)
}

func main() {
	router := gin.Default()

	// Configure trusted proxies
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"}) // Replace with your trusted proxy IP or IPs

	// Register helloHandler
	router.GET("/hello", helloHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running at port %s\n", port)
	err := router.Run(":" + port)
	if err != nil {
		// Ada kesalahan saat menjalankan server
		panic(fmt.Sprintf("Failed to start server: %v", err))
	} else {
		fmt.Printf("Server running at port %s\n", port)
	}

}
