package userRepo

import (
	"context"
	"time"

	interfaces "github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/_interfaces"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/video-users/src/repositories/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) interfaces.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (ref *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	record := models.UserFromDomain(user)

	now := time.Now()
	record.ID = uuid.NewString()
	record.CreatedAt = now
	record.UpdatedAt = now

	_, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": record.ID})
	if err := result.Err(); err != nil {
		return nil, err
	}

	var foundUser models.User
	if err := result.Decode(&foundUser); err != nil {
		return nil, err
	}

	return foundUser.ToDomain(), nil
}

func (ref *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := ref.collection.FindOne(ctx, bson.M{"email": email})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, result.Err()
	}

	var foundUser models.User
	if err := result.Decode(&foundUser); err != nil {
		return nil, err
	}

	return foundUser.ToDomain(), nil
}
