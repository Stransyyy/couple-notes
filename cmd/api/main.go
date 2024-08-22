package main

import (
	"github.com/stransyyy/couple-notes/config"
	"github.com/stransyyy/couple-notes/internal/handlers"
	"github.com/stransyyy/couple-notes/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()

	config.LoadConfig()
	config.InitDynamoDB()

	store := &repository.DynamoDBStore{DB: config.DynamoDB}

	r := gin.Default()

	// Define routes and map them to the handlers
	r.POST("/notes", handlers.CreateNote(store))
	r.GET("/notes/:id", handlers.GetNoteByID(store))
	r.PUT("/notes/:id", handlers.UpdateNoteContent(store))
	r.DELETE("/notes/:id", handlers.DeleteNoteByID(store))

	r.Run(":8080")
}
