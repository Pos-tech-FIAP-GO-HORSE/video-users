package userService

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	mocks "github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/_mocks"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	jwtSecretKey := uuid.NewString()
	expectedError := errors.New("unexpected error")
	now := time.Now()

	t.Run("should not create user when failed o find user by email", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(nil, expectedError)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should not create user when email already exists", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(&user, nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, errors.New("email already exists"), err)
		assert.Nil(t, actual)
	})

	t.Run("should not create user when failed to generate password hash", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(nil, nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		user.Password = string(make([]byte, 73))

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, errors.New("bcrypt: password length exceeds 72 bytes"), err)
		assert.Nil(t, actual)
	})

	t.Run("should not create user when failed to create user in database", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(nil, nil)
		userRepositoryMocked.On("Create", ctx, &user).Return(nil, expectedError)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should not create user when failed to publish the message to broker", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(nil, nil)
		userRepositoryMocked.On("Create", ctx, &user).Return(&user, nil)

		messageRaw, _ := json.Marshal(user)
		publisherMocked := mocks.NewPublisher(t)
		publisherMocked.On("Publish", ctx, string(messageRaw)).Return(expectedError)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, publisherMocked)

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should create user successfully", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        "some@email.com",
			Password:     uuid.NewString(),
			PasswordHash: uuid.NewString(),
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, user.Email).Return(nil, nil)
		userRepositoryMocked.On("Create", ctx, &user).Return(&user, nil)

		messageRaw, _ := json.Marshal(user)
		publisherMocked := mocks.NewPublisher(t)
		publisherMocked.On("Publish", ctx, string(messageRaw)).Return(nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, publisherMocked)

		expected := &user

		actual, err := userService.Create(ctx, &user)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	ctx := context.TODO()
	jwtSecretKey := uuid.NewString()
	email := "some@email.com"
	password := "some-password"
	now := time.Now()
	expectedError := errors.New("unexpected error")

	t.Run("should not login when failed to find user by email", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, email).Return(nil, expectedError)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Login(ctx, email, password)
		assert.Equal(t, "", actual)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should not login when email does not exist", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, email).Return(nil, nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Login(ctx, email, password)
		assert.Equal(t, "", actual)
		assert.Equal(t, errors.New("email does not exist"), err)
	})

	t.Run("should not login when password does not match", func(t *testing.T) {
		user := entity.User{
			ID:           uuid.NewString(),
			Name:         "Some Name",
			Email:        email,
			Password:     "other-password",
			PasswordHash: "other-password",
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, email).Return(&user, nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Login(ctx, email, password)
		assert.Equal(t, "", actual)
		assert.Equal(t, errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password"), err)
	})

	t.Run("should not login when password does not match", func(t *testing.T) {
		user := entity.User{
			ID:        uuid.NewString(),
			Name:      "Some Name",
			Email:     email,
			Password:  "other-password",
			CreatedAt: now,
			UpdatedAt: now,
		}

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.PasswordHash = string(passwordHash)

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("FindByEmail", ctx, email).Return(&user, nil)

		userService := NewUserService(jwtSecretKey, userRepositoryMocked, nil)

		actual, err := userService.Login(ctx, email, password)
		assert.Equal(t, "", actual)
		assert.Equal(t, errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"), err)
	})
}
