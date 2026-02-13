package dto

import (
	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func ConvertToEntityLoginRequest(request LoginUserRequest) entities.User {
	return entities.User{
		Username: request.Username,
		Password: request.Password,
	}
}

func ConvertToEntityUserRequest(request RegisterUserRequest) entities.User {
	return entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}
}

func (u *RegisterUserRequest) HashPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "User.HashPassword.GenerateFromPassword")
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *RegisterUserRequest) PrepareCreate() error {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}
func (u *LoginUserRequest) ComparePassword(hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password)); err != nil {
		return errors.Wrap(err, "User.ComparePassword.CompareHashAndPassword")
	}
	return nil
}

func (u *UpdateUserRequest) HashPassword() error {
	if u.Password == "" {
		return nil // jika kosong, tidak perlu hash
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "User.Update.HashPassword.GenerateFromPassword")
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *UpdateUserRequest) PrepareUpdate() error {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}
