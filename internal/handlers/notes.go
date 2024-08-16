package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stransyyy/couple-notes/internal/models"
	"github.com/stransyyy/couple-notes/internal/repository"
)

// CreateNote handles the POST request to add a new note to the DB
func CreateNote(store repository.DynamoDBStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var note models.Note
		if err := ctx.BindJSON(&note); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Input"})
			return
		}

		note.ID = uuid.New().String()
		note.CreatedAt = time.Now()

		err := store.SaveNote(note)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to save note on the database"})
			return
		}

		ctx.JSON(200, gin.H{"error": "Note saved successfully", "note": note})
	}
}

// GetNote handles the GET request to fetch a note by ID.
func GetNoteByID(store repository.DynamoDBStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		note, err := store.GetNoteByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Note not found"})
			return
		}

		c.JSON(200, gin.H{"note": note})
	}
}

// UpdateNoteContent handles the PUT request to update a note's content.
func UpdateNoteContent(store repository.DynamoDBStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payload models.Note

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Call the repository and pass the ID separately from the Note object
		err := store.UpdateNoteContent(id, payload)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update note"})
			return
		}

		c.JSON(200, gin.H{"message": "Note updated successfully"})
	}
}

// DeleteNote handles the DELETE request to remove a note by ID.
func DeleteNoteByID(store repository.DynamoDBStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := store.DeleteNoteByID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete note"})
			return
		}

		c.JSON(200, gin.H{"message": "Note deleted successfully"})
	}
}
