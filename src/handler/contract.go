package handler

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c CreateUserRequest) ToDomain() *entity.User {
	return &entity.User{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
