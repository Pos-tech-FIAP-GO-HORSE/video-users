package userService

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	interfaces "github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/_interfaces"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	jwtSecretKey   string
	userRepository interfaces.UserRepository
	publisher      interfaces.Publisher
}

func NewUserService(jwtSecretKey string, userRepository interfaces.UserRepository, publisher interfaces.Publisher) interfaces.UserService {
	return &userService{
		jwtSecretKey:   jwtSecretKey,
		userRepository: userRepository,
		publisher:      publisher,
	}
}

func (ref *userService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
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

	createdUser, err := ref.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	messageRaw, _ := json.Marshal(createdUser)
	if err = ref.publisher.Publish(ctx, string(messageRaw)); err != nil {
		return nil, err
	}

	log.Println("[INFO] published message:", string(messageRaw))

	return createdUser, nil
}

func (ref *userService) Login(ctx context.Context, email string, password string) (string, error) {
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

	tokenString, err := token.SignedString([]byte(ref.jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
