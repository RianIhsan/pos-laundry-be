package repository

import (
	"context"
	"fmt"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type servicePostgresRepository struct {
	db *gorm.DB
}

func NewServicePostgresRepository(db *gorm.DB) services.ServiceRepositoryInterface {
	return &servicePostgresRepository{db: db}
}

func (s *servicePostgresRepository) Create(ctx context.Context, entity entities.Service) (entities.Service, error) {
	if err := s.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return entities.Service{}, err
	}
	return entity, nil
}

func (s *servicePostgresRepository) GetList(ctx context.Context) ([]entities.Service, error) {
	var services []entities.Service
	query := s.db.WithContext(ctx).Model(&entities.Service{})

	if err := query.Order("created_at ASC").Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (s *servicePostgresRepository) FindById(ctx context.Context, serviceId uint) (entities.Service, error) {
	var service entities.Service
	DB := s.db.WithContext(ctx)

	if err := DB.Where("id = ?", serviceId).Take(&service).Error; err != nil {
		return entities.Service{}, err
	}
	return service, nil
}

func (s *servicePostgresRepository) Update(ctx context.Context, id uint, data entities.Service) error {
	var existing entities.Service
	if err := s.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("service with id %d not found", id)
		}
		return fmt.Errorf("cannot update service with id %d: %w", id, err)
	}
	if err := s.db.WithContext(ctx).Model(&existing).Updates(data).Error; err != nil {
		return fmt.Errorf("cannot update service with id %d: %w", id, err)
	}
	return nil
}

func (s *servicePostgresRepository) DeleteService(ctx context.Context, serviceId uint) error {
	if err := s.db.WithContext(ctx).Delete(&entities.Service{}, serviceId).Error; err != nil {
		return err
	}
	return nil
}
