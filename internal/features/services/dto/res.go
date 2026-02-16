package dto

import (
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type ServiceResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
}

type ServiceDTO = ServiceResponse

func ToServiceDTO(e entities.Service) ServiceDTO {
	return ServiceDTO{
		ID:        e.ID,
		Name:      e.Name,
		Category:  e.Category,
		Price:     e.Price,
		Unit:      e.Unit,
		CreatedAt: e.CreatedAt,
	}
}

func ToListServicesResponse(list []entities.Service) (resp []ServiceDTO) {
	for _, s := range list {
		resp = append(resp, ToServiceDTO(s))
	}
	return resp
}
