package board

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *boardRepository) GetList(customerId string) ([]entity.Board, error) {
	var models model.BoardList

	err := r.db.Model(&models).Where("customer_id = ?", uuid.MustParse(customerId)).Select()
	if err != nil {
		return nil, errors.Wrap(err, "select boards list error")
	}

	return models.BoardListEntity(), nil
}
