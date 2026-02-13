package repository

import (
	"context"
	"fmt"
	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) users.UserRepositoryInterface {
	return &userPostgresRepository{db: db}
}

func (u *userPostgresRepository) Create(ctx context.Context, entity entities.User) (entities.User, error) {
	if err := u.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return entities.User{}, err
	}
	return entity, nil
}

func (u *userPostgresRepository) FindByUsername(ctx context.Context, entity entities.User) (*entities.User, error) {
	user := new(entities.User)
	DB := u.db.WithContext(ctx)

	if err := DB.Where(entities.User{Username: entity.Username}).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPostgresRepository) GetList(ctx context.Context, roleId uint64) ([]entities.User, error) {
	var users []entities.User
	query := u.db.WithContext(ctx).Model(&entities.User{}).
		Preload("Role")
	//Preload("Branch").
	//Preload("Company")

	if roleId != 0 {
		query = query.Where("role_id = ?", roleId)
	}

	if err := query.Order("created_at ASC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userPostgresRepository) FindById(ctx context.Context, userId uint64) (entities.User, error) {
	var user entities.User
	DB := u.db.WithContext(ctx)
	//Preload("Branch").
	//Preload("Company")

	if err := DB.Where("id = ?", userId).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *userPostgresRepository) Update(ctx context.Context, id uint64, data entities.User) error {
	var existingUser entities.User
	if err := u.db.WithContext(ctx).First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %d not found", id)
		}
		return fmt.Errorf("cannot update user with id %d: %w", id, err)
	}
	if err := u.db.WithContext(ctx).Model(&existingUser).Updates(data).Error; err != nil {
		return fmt.Errorf("cannot update user with id %d: %w", id, err)
	}
	return nil
}

func (u *userPostgresRepository) DeleteUser(ctx context.Context, userId uint64) error {
	if err := u.db.WithContext(ctx).Delete(&entities.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (u *userPostgresRepository) UpdatePassword(ctx context.Context, email, hashedPassword string) error {
	return u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("email = ?", email).
		Update("password", hashedPassword).Error
}
