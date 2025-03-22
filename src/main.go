package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/useCases/users_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/handler"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/repositories/user_repository"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("unable to connect on database: %v", err)
	}

	log.Println("database connected successfully")

	database := dbClient.Database(os.Getenv("DB_NAME"))
	usersCollection := database.Collection("users")

	userRepository := user_repository.NewUserRepository(usersCollection)
	userService := users_service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	lambda.Start(userHandler.Handle)
}
