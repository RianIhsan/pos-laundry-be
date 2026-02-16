package service

import (
	"context"
	"fmt"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers/dto"
	"github.com/pkg/errors"
)

type customerService struct {
	cfg          *ServiceConfig
	customerRepo customers.CustomerRepositoryInterface
}

func NewCustomerService(cfg *ServiceConfig) customers.CustomerServiceInterface {
	return &customerService{
		cfg:          cfg,
		customerRepo: cfg.CustomerRepoInterface,
	}
}

func (s *customerService) AddCustomer(ctx context.Context, userID uint, req dto.CreateCustomerRequest) (dto.CustomerResponse, error) {
	req.PrepareCreate()

	created, err := s.customerRepo.Create(ctx, dto.ConvertToEntityCustomerRequest(req))
	if err != nil {
		return dto.CustomerResponse{}, errors.Wrap(err, "failed to create customer")
	}

	// Log activity
	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogCustomerCreated(
			ctx,
			userID,
			fmt.Sprintf("%d", created.ID),
			created.Name,
		)
	}

	return dto.ToCustomerDTO(created), nil
}

func (s *customerService) GetList(ctx context.Context) ([]dto.CustomerDTO, error) {
	list, err := s.customerRepo.GetList(ctx)
	if err != nil {
		return []dto.CustomerDTO{}, errors.Wrap(err, "failed to list customers")
	}
	return dto.ToListCustomersResponse(list), nil
}

func (s *customerService) GetById(ctx context.Context, customerId uint) (dto.CustomerDTO, error) {
	found, err := s.customerRepo.FindById(ctx, customerId)
	if err != nil {
		return dto.CustomerDTO{}, errors.Wrap(err, "failed to find customer")
	}
	return dto.ToCustomerDTO(found), nil
}

func (s *customerService) Update(ctx context.Context, userID uint, id uint, data dto.UpdateCustomerRequest) error {
	data.PrepareUpdate()

	// Get customer before update to log
	customer, err := s.customerRepo.FindById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to find customer")
	}

	err = s.customerRepo.Update(ctx, id, entities.Customer{
		Name:    data.Name,
		Phone:   data.Phone,
		Address: data.Address,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update customer")
	}

	// Log activity
	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogCustomerUpdated(
			ctx,
			userID,
			fmt.Sprintf("%d", id),
			customer.Name,
		)
	}

	return nil
}

func (s *customerService) Delete(ctx context.Context, userID uint, customerId uint) error {
	customer, err := s.customerRepo.FindById(ctx, customerId)
	if err != nil {
		return errors.Wrap(err, "failed to find customer")
	}

	if err := s.customerRepo.DeleteCustomer(ctx, customerId); err != nil {
		return errors.Wrap(err, "failed to delete customer")
	}

	// Log activity
	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogCustomerDeleted(
			ctx,
			userID,
			fmt.Sprintf("%d", customerId),
			customer.Name,
		)
	}

	return nil
}
