package customer

import (
	"kanban-bot/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type customerRepository struct {
	db orm.DB
}

func NewCustomerRepository(db orm.DB) interfaces.CustomerRepository {
	return &customerRepository{db: db}
}
