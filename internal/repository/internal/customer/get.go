package customer

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *customerRepository) Get(id string) (entity.Customer, error) {
	model := model.NewCustomer(entity.Customer{ID: id})

	err := r.db.Model(&model).WherePK().Select()
	if err != nil {
		return entity.Customer{}, errors.Wrap(err, "select customer error")
	}

	return model.CustomerEntity(), nil
}
