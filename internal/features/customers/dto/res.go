package dto

import (
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type CustomerResponse struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	Points  int       `json:"points"`
	Created time.Time `json:"created_at"`
}

type CustomerDTO = CustomerResponse

func ToCustomerDTO(e entities.Customer) CustomerDTO {
	return CustomerDTO{
		ID:      e.ID,
		Name:    e.Name,
		Phone:   e.Phone,
		Address: e.Address,
		Points:  e.Points,
		Created: e.CreatedAt,
	}
}

func ToListCustomersResponse(list []entities.Customer) (resp []CustomerDTO) {
	for _, c := range list {
		resp = append(resp, ToCustomerDTO(c))
	}
	return resp
}
