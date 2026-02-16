package services

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/features/services/dto"
)

type ServiceServiceInterface interface {
	AddService(ctx context.Context, req dto.CreateServiceRequest) (dto.ServiceResponse, error)
	GetList(ctx context.Context) ([]dto.ServiceDTO, error)
	GetById(ctx context.Context, serviceId uint) (dto.ServiceDTO, error)
	Update(ctx context.Context, id uint, data dto.UpdateServiceRequest) error
	Delete(ctx context.Context, serviceId uint) error
}
