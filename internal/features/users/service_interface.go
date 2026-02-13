package users

import (
	"context"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users/dto"
)

type UserServiceInterface interface {
	AddUser(ctx context.Context, request dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	LoginUser(ctx context.Context, request *dto.LoginUserRequest) (*dto.JwtToken, error)
	GetList(ctx context.Context, roleId uint64) ([]dto.UserDTO, error)
	GetById(ctx context.Context, userId uint64) (dto.UserDTO, error)
	Delete(ctx context.Context, userId uint64) error
	Update(ctx context.Context, id uint64, data dto.UpdateUserRequest) error
}
