package dto

import (
	"strings"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required,min=1,max=100"`
	Phone   string `json:"phone" validate:"required,min=3,max=20"`
	Address string `json:"address"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (r *CreateCustomerRequest) PrepareCreate() {
	r.Name = strings.TrimSpace(r.Name)
	r.Phone = strings.TrimSpace(r.Phone)
	r.Address = strings.TrimSpace(r.Address)
}

func (r *UpdateCustomerRequest) PrepareUpdate() {
	r.Name = strings.TrimSpace(r.Name)
	r.Phone = strings.TrimSpace(r.Phone)
	r.Address = strings.TrimSpace(r.Address)
}

func ConvertToEntityCustomerRequest(req CreateCustomerRequest) entities.Customer {
	return entities.Customer{
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
}
