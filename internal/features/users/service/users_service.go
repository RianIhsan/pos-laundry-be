package service

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users/dto"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/pkg/errors"
)

type userService struct {
	cfg      *config.Config
	userRepo users.UserRepositoryInterface
}

func NewUserService(cfg *ServiceConfig) users.UserServiceInterface {
	return &userService{
		cfg:      cfg.Config,
		userRepo: cfg.UserRepoInterface,
	}
}

func (uS *userService) AddUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	result, err := uS.userRepo.FindByUsername(ctx, entities.User{Username: req.Username})
	if result != nil && err == nil {
		return dto.RegisterUserResponse{}, errors.New("Email already exist")
	}

	if err := req.PrepareCreate(); err != nil {
		return dto.RegisterUserResponse{}, errors.Wrap(err, "failed to prepare user data")
	}

	createdUser, err := uS.userRepo.Create(ctx, dto.ConvertToEntityUserRequest(req))
	if err != nil {
		return dto.RegisterUserResponse{}, errors.Wrap(err, "failed to create user")
	}

	return dto.ConvertToRegisterUserResponse(&createdUser), nil
}

func (uS *userService) LoginUser(ctx context.Context, request *dto.LoginUserRequest) (*dto.JwtToken, error) {
	foundUser, err := uS.userRepo.FindByUsername(ctx, entities.User{Username: request.Username})
	if err != nil {
		return nil, errors.Wrap(err, "username not found")
	}

	if err := request.ComparePassword(foundUser.Password); err != nil {
		return nil, errors.New("invalid password")
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(foundUser, uS.cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate tokens")
	}

	return &dto.JwtToken{
		Username:     foundUser.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role:         foundUser.Role,
	}, nil
}

func (uS *userService) GetList(ctx context.Context, roleId uint64) ([]dto.UserDTO, error) {
	data, err := uS.userRepo.GetList(ctx, roleId)
	if err != nil {
		return []dto.UserDTO{}, errors.Wrap(err, "failed to list users")
	}
	return dto.ToListUsersResponse(data), nil
}

func (uS *userService) GetById(ctx context.Context, userId uint64) (dto.UserDTO, error) {
	fetchUser, err := uS.userRepo.FindById(ctx, userId)
	if err != nil {
		return dto.UserDTO{}, errors.Wrap(err, "failed to find user")
	}

	return dto.ToUserDTO(fetchUser), nil
}

func (uS *userService) Update(ctx context.Context, id uint64, data dto.UpdateUserRequest) error {
	if err := data.PrepareUpdate(); err != nil {
		return errors.Wrap(err, "failed to prepare update user")
	}

	err := uS.userRepo.Update(ctx, id, entities.User{
		Username: data.Username,
		Name:     data.Name,
		Password: data.Password,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (uS *userService) Delete(ctx context.Context, userId uint64) error {
	_, err := uS.userRepo.FindById(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	if err := uS.userRepo.DeleteUser(ctx, userId); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}
