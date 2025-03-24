package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/useCases/users_service"
	"github.com/aws/aws-lambda-go/events"
)

type UserHandler struct {
	userService users_service.IUserService
}

func NewUserHandler(userService users_service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (ref *UserHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
	case request.Path == "/users" && request.HTTPMethod == http.MethodPost:
		return ref.handleCreateUser(ctx, request)
	case request.Path == "/users/login" && request.HTTPMethod == http.MethodPost:
		return ref.handleLogin(ctx, request)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       "Not Found",
	}, nil
}

func (ref *UserHandler) handleCreateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body CreateUserRequest
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("error to decode create request body:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid JSON: " + err.Error()}, nil
	}

	user, err := ref.userService.Create(ctx, body.ToDomain())
	if err != nil {
		log.Println("error to create user:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Internal Server Error: " + err.Error()}, nil
	}

	response := CreateUserResponseFromDomain(user)
	responseRaw, _ := json.Marshal(response)

	log.Println("user created successfully")
	return events.APIGatewayProxyResponse{StatusCode: http.StatusCreated, Body: string(responseRaw)}, nil
}

func (ref *UserHandler) handleLogin(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body LoginRequest
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("error to decode login request body:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid JSON: " + err.Error()}, nil
	}

	accessToken, err := ref.userService.Login(ctx, body.Email, body.Password)
	if err != nil {
		log.Println("error to login:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized: " + err.Error()}, nil
	}

	response := LoginResponse{AccessToken: accessToken}
	responseRaw, _ := json.Marshal(response)

	log.Println("user logged in successfully")
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(responseRaw)}, nil
}
