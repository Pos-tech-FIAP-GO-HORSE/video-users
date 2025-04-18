package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/responses"
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

func (ref *UserHandler) Handle(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch {
	case request.RawPath == "/video-user/users" && request.RequestContext.HTTP.Method == http.MethodPost:
		return ref.handleCreateUser(ctx, request)
	case request.RawPath == "/video-user/users/login" && request.RequestContext.HTTP.Method == http.MethodPost:
		return ref.handleLogin(ctx, request)
	}

	return apiGatewayResponseWithError(http.StatusNotFound, "route not found"), nil
}

func (ref *UserHandler) handleCreateUser(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body CreateUserRequest
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("[ERROR] error to decode create request body:", err)
		return apiGatewayResponseWithError(http.StatusBadRequest, "invalid json: "+err.Error()), nil
	}

	user, err := ref.userService.Create(ctx, body.ToDomain())
	if err != nil {
		log.Println("[ERROR] error to create user:", err)
		return apiGatewayResponseWithError(http.StatusInternalServerError, err.Error()), nil
	}

	response := responses.UserFromDomain(user)
	responseRaw, _ := json.Marshal(response)

	log.Println("[INFO] user created successfully")
	return apiGatewayResponse(http.StatusCreated, string(responseRaw)), nil
}

func (ref *UserHandler) handleLogin(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body LoginRequest
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("[ERROR] error to decode login request body:", err)
		return apiGatewayResponseWithError(http.StatusBadRequest, "invalid json: "+err.Error()), nil
	}

	accessToken, err := ref.userService.Login(ctx, body.Email, body.Password)
	if err != nil {
		log.Println("[ERROR] error to login:", err)
		return apiGatewayResponseWithError(http.StatusUnauthorized, "unauthorized: "+err.Error()), nil
	}

	response := responses.LoginResponse{AccessToken: accessToken}
	responseRaw, _ := json.Marshal(response)

	log.Println("[INFO] user logged in successfully")
	return apiGatewayResponse(http.StatusOK, string(responseRaw)), nil
}
