package model

import (
	"kanban-bot/internal/core/entity"

	"github.com/google/uuid"
)

type Customer struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

func NewCustomer(entity entity.Customer) Customer {
	return Customer{
		ID:       uuid.MustParse(entity.ID),
		Username: entity.Username,
	}
}

func (c *Customer) CustomerEntity() entity.Customer {
	return entity.Customer{
		ID:       c.ID.String(),
		Username: c.Username,
	}
}
