package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/useCases/userService"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/handler"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/publisher"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/repositories/userRepo"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	var (
		// Database
		dbURI  = os.Getenv("DB_URI")
		dbName = os.Getenv("DB_NAME")

		// JWT
		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

		// SNS
		userSnsTopic = os.Getenv("USER_TOPIC_ARN")
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbURI)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("[ERROR] unable to connect on database: %v", err)
	}

	if err = dbClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("[ERROR] unable to ping the database: %v", err)
	}

	database := dbClient.Database(dbName)
	usersCollection := database.Collection("users")

	log.Println("[INFO] database connected successfully")

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		log.Fatalf("[ERROR] unable to load aws config: %v", err)
	}

	userPublisher := publisher.NewSnsPublisher(sns.NewFromConfig(cfg), userSnsTopic)

	userRepository := userRepo.NewUserRepository(usersCollection)
	userService := userService.NewUserService(jwtSecretKey, userRepository, userPublisher)
	userHandler := handler.NewUserHandler(userService)

	lambda.Start(userHandler.Handle)
}
