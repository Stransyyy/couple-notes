package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var DynamoDB *dynamodb.Client
var Log *logrus.Logger

// InitLogger initializes the Logrus logger.
func InitLogger() {
	if Log == nil {
		Log = logrus.New()
		Log.Out = os.Stdout
		Log.SetLevel(logrus.InfoLevel)
		log.Println("Logger initialized successfully")
	} else {
		log.Println("Logger is already initialized")
	}

}

// InitDynamoDB initializes the DynamoDB connection.
func InitDynamoDB() {
	creds := credentials.NewStaticCredentialsProvider(*aws.String(), *aws.String(), "")

	// Load AWS SDK configuration using the default method.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds),
		config.WithRegion(*aws.String(os.Getenv("AWS_REGION"))))
	if err != nil {
		Log.Info("Unable to load the AWS SDK Config", err.Error())
	}

	// Initialize DynamoDB client.
	DynamoDB = dynamodb.NewFromConfig(cfg)
	Log.Info("Connected to DynamoDB successfully")
}

// LoadConfig loads any additional configurations, such as environment variables.
func LoadConfig() {
	// If using an .env file, you can load it here using a library like godotenv (optional).
	// err := godotenv.Load()
	// if err != nil {
	//     log.Fatal("Error loading .env file")
	// }

	Log.Info("Starting")
}
