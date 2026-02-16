package services

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type ServiceRepositoryInterface interface {
	Create(ctx context.Context, service entities.Service) (entities.Service, error)
	GetList(ctx context.Context) ([]entities.Service, error)
	FindById(ctx context.Context, serviceId uint) (entities.Service, error)
	Update(ctx context.Context, id uint, data entities.Service) error
	DeleteService(ctx context.Context, serviceId uint) error
}
