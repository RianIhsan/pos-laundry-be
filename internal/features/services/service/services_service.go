package service

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services/dto"
	"github.com/pkg/errors"
)

type serviceService struct {
	cfg                  *ServiceConfig
	serviceRepoInterface services.ServiceRepositoryInterface
}

func NewServiceService(cfg *ServiceConfig) services.ServiceServiceInterface {
	return &serviceService{
		cfg:                  cfg,
		serviceRepoInterface: cfg.ServiceRepoInterface,
	}
}

func (s *serviceService) AddService(ctx context.Context, req dto.CreateServiceRequest) (dto.ServiceResponse, error) {
	req.PrepareCreate()

	created, err := s.serviceRepoInterface.Create(ctx, dto.ConvertToEntityServiceRequest(req))
	if err != nil {
		return dto.ServiceResponse{}, errors.Wrap(err, "failed to create service")
	}
	return dto.ToServiceDTO(created), nil
}

func (s *serviceService) GetList(ctx context.Context) ([]dto.ServiceDTO, error) {
	list, err := s.serviceRepoInterface.GetList(ctx)
	if err != nil {
		return []dto.ServiceDTO{}, errors.Wrap(err, "failed to list services")
	}
	return dto.ToListServicesResponse(list), nil
}

func (s *serviceService) GetById(ctx context.Context, serviceId uint) (dto.ServiceDTO, error) {
	found, err := s.serviceRepoInterface.FindById(ctx, serviceId)
	if err != nil {
		return dto.ServiceDTO{}, errors.Wrap(err, "failed to find service")
	}
	return dto.ToServiceDTO(found), nil
}

func (s *serviceService) Update(ctx context.Context, id uint, data dto.UpdateServiceRequest) error {
	data.PrepareUpdate()

	err := s.serviceRepoInterface.Update(ctx, id, entities.Service{
		Name:     data.Name,
		Category: data.Category,
		Price:    data.Price,
		Unit:     data.Unit,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update service")
	}
	return nil
}

func (s *serviceService) Delete(ctx context.Context, serviceId uint) error {
	_, err := s.serviceRepoInterface.FindById(ctx, serviceId)
	if err != nil {
		return errors.Wrap(err, "failed to find service")
	}

	if err := s.serviceRepoInterface.DeleteService(ctx, serviceId); err != nil {
		return errors.Wrap(err, "failed to delete service")
	}
	return nil
}
