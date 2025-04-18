package interfaces

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
