package models

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type User struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	IntegrationID string    `json:"integrationId" bson:"integrationId"`
	Name          string    `json:"name" bson:"name"`
	Email         string    `json:"email" bson:"email"`
	PasswordHash  string    `json:"passwordHash" bson:"passwordHash"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (ref User) ToDomain() *entity.User {
	return &entity.User{
		ID:            ref.ID,
		IntegrationID: ref.IntegrationID,
		Name:          ref.Name,
		Email:         ref.Email,
		PasswordHash:  ref.PasswordHash,
		CreatedAt:     ref.CreatedAt,
		UpdatedAt:     ref.UpdatedAt,
	}
}

func UserFromDomain(user *entity.User) User {
	return User{
		ID:            user.ID,
		IntegrationID: user.IntegrationID,
		Name:          user.Name,
		Email:         user.Email,
		PasswordHash:  user.PasswordHash,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
