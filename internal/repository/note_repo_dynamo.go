package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stransyyy/couple-notes/config"
	"github.com/stransyyy/couple-notes/internal/models"
)

type DynamoDBStore struct {
	DB *dynamodb.Client
}

func (d *DynamoDBStore) SaveNote(note models.Note) error {
	// Convert time.Time to a string
	createdAt := note.CreatedAt.Format(time.RFC3339)

	// Prepare the item to insert into DynamoDB
	item := map[string]types.AttributeValue{
		"NoteID":    &types.AttributeValueMemberS{Value: note.ID},
		"Content":   &types.AttributeValueMemberS{Value: note.Content},
		"UserID":    &types.AttributeValueMemberS{Value: note.UserID},
		"CreatedAt": &types.AttributeValueMemberS{Value: createdAt},
	}

	// Put item in the DynamoDB table
	_, err := d.DB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("NotesTable"), // Use your actual table name here
		Item:      item,
	})

	if err != nil {
		config.Log.Error("Failed to save note to DynamoDB", err)
		return err
	}

	config.Log.Info("Note saved successfully to DynamoDB", note.ID)
	return nil
}

func (d *DynamoDBStore) GetNoteByID(id string) (models.Note, error) {
	result, err := d.DB.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("NotesTable"),
		Key: map[string]types.AttributeValue{
			"NoteID": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		config.Log.Error("Failed to retrieve note from DynamoDB", err)
		return models.Note{}, err
	}

	// Parse CreatedAt string back to time.Time
	createdAt, err := time.Parse(time.RFC3339, result.Item["CreatedAt"].(*types.AttributeValueMemberS).Value)
	if err != nil {
		return models.Note{}, err
	}

	note := models.Note{
		ID:        id,
		Content:   result.Item["Content"].(*types.AttributeValueMemberS).Value,
		UserID:    result.Item["UserID"].(*types.AttributeValueMemberS).Value,
		CreatedAt: createdAt,
	}

	return note, nil
}
