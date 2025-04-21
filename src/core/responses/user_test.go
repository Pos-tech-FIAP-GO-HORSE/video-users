package responses

import (
	"testing"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserFromDomain(t *testing.T) {
	now := time.Now()
	userID := uuid.NewString()
	name := "Some name"
	email := "some@email.com"
	password := "some-password"
	passwordHash := "some-password-hash"

	user := UserFromDomain(&entity.User{
		ID:           userID,
		Name:         name,
		Email:        email,
		Password:     password,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	assert.Equal(t, userID, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}
