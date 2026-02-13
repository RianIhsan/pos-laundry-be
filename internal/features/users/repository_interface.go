package users

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user entities.User) (entities.User, error)
	FindByUsername(ctx context.Context, user entities.User) (*entities.User, error)
	GetList(ctx context.Context, roleId uint64) ([]entities.User, error)
	FindById(ctx context.Context, userId uint64) (entities.User, error)
	Update(ctx context.Context, id uint64, data entities.User) error
	DeleteUser(ctx context.Context, userId uint64) error
}
