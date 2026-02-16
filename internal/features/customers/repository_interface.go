package customers

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type CustomerRepositoryInterface interface {
	Create(ctx context.Context, customer entities.Customer) (entities.Customer, error)
	GetList(ctx context.Context) ([]entities.Customer, error)
	FindById(ctx context.Context, customerId uint) (entities.Customer, error)
	Update(ctx context.Context, id uint, data entities.Customer) error
	DeleteCustomer(ctx context.Context, customerId uint) error
}
