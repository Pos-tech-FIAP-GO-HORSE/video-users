package models

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type User struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Name         string    `json:"name" bson:"name"`
	Email        string    `json:"email" bson:"email"`
	PasswordHash string    `json:"password_hash" bson:"password_hash"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func (ref User) ToDomain() *entity.User {
	return &entity.User{
		ID:           ref.ID,
		Name:         ref.Name,
		Email:        ref.Email,
		PasswordHash: ref.PasswordHash,
		CreatedAt:    ref.CreatedAt,
		UpdatedAt:    ref.UpdatedAt,
	}
}

func UserFromDomain(user *entity.User) User {
	return User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
