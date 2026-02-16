package customers

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/features/customers/dto"
)

type CustomerServiceInterface interface {
	AddCustomer(ctx context.Context, userID uint, req dto.CreateCustomerRequest) (dto.CustomerResponse, error)
	GetList(ctx context.Context) ([]dto.CustomerDTO, error)
	GetById(ctx context.Context, customerId uint) (dto.CustomerDTO, error)
	Update(ctx context.Context, userID uint, id uint, data dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, userID uint, customerId uint) error
}
