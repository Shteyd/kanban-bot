package customer

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *customerRepository) Save(customer *entity.Customer) error {
	model := model.NewCustomer(*customer)
	if customer.ID == "" {
		_, err := r.db.Model(&model).Insert()
		if err != nil {
			return errors.Wrap(err, "insert customer error")
		}
	} else {
		_, err := r.db.Model(&model).WherePK().Update()
		if err != nil {
			return errors.Wrap(err, "update customer error")
		}
	}

	return nil
}
