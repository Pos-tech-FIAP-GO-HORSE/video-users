package responses

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type User struct {
	ID            string    `json:"id"`
	IntegrationID string    `json:"integration_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func UserFromDomain(user *entity.User) User {
	return User{
		ID:            user.ID,
		IntegrationID: user.IntegrationID,
		Name:          user.Name,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
