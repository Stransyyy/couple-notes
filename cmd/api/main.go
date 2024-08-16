package main

import (
	"fmt"

	"github.com/stransyyy/couple-notes/config"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Initializing logger...")
	config.InitLogger() // Initialize the logger first

	fmt.Println("Loading config...")
	config.LoadConfig() // Load configuration after initializing the logger

	fmt.Println("Initializing DynamoDB...")
	config.InitDynamoDB() // Initialize DynamoDB connection

	fmt.Println("Starting Gin server...")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		config.Log.Info("Ping received")
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")

}
