package dto

import (
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type RegisterUserResponse struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func ConvertToRegisterUserResponse(entity *entities.User) RegisterUserResponse {
	return RegisterUserResponse{
		Id:       entity.ID,
		Name:     entity.Name,
		Username: entity.Username,
		Role:     entity.Role,
	}
}

func ToListUsers(users []entities.User) (response []RegisterUserResponse) {
	for _, user := range users {
		response = append(response, ConvertToRegisterUserResponse(&user))
	}
	return response
}

type JwtToken struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

type UserDTO struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserDTO(user entities.User) UserDTO {

	return UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

func ToListUsersResponse(users []entities.User) (response []UserDTO) {
	for _, user := range users {
		response = append(response, ToUserDTO(user))
	}
	return response
}
