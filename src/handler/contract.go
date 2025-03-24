package handler

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID            string    `json:"id"`
	IntegrationID string    `json:"integrationId"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (c CreateUserRequest) ToDomain() *entity.User {
	return &entity.User{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}

func CreateUserResponseFromDomain(user *entity.User) CreateUserResponse {
	return CreateUserResponse{
		ID:            user.ID,
		IntegrationID: user.IntegrationID,
		Name:          user.Name,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
