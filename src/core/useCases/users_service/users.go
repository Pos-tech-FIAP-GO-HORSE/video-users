package users_service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/repositories/user_repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type UserService struct {
	userRepository user_repository.IUserRepository
}

func NewUserService(userRepository user_repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (ref *UserService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	foundUser, err := ref.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if foundUser != nil {
		return nil, errors.New("email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = string(passwordHash)

	return ref.userRepository.Create(ctx, user)
}

func (ref *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	foundUser, err := ref.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if foundUser == nil {
		return "", errors.New("email does not exist")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": foundUser.ID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
