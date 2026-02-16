package dto

import (
	"strings"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type CreateServiceRequest struct {
	Name     string  `json:"name" validate:"required,min=1,max=100"`
	Category string  `json:"category" validate:"required,min=1,max=50"`
	Price    float64 `json:"price" validate:"required,min=0"`
	Unit     string  `json:"unit" validate:"required,min=1,max=10"`
}

type UpdateServiceRequest struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Unit     string  `json:"unit"`
}

func (r *CreateServiceRequest) PrepareCreate() {
	r.Name = strings.TrimSpace(r.Name)
	r.Category = strings.TrimSpace(r.Category)
	r.Unit = strings.TrimSpace(r.Unit)
}

func (r *UpdateServiceRequest) PrepareUpdate() {
	r.Name = strings.TrimSpace(r.Name)
	r.Category = strings.TrimSpace(r.Category)
	r.Unit = strings.TrimSpace(r.Unit)
}

func ConvertToEntityServiceRequest(req CreateServiceRequest) entities.Service {
	return entities.Service{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Unit:     req.Unit,
	}
}
