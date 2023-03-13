package board

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *boardRepository) Save(board *entity.Board) error {
	model := model.NewBoard(*board)

	if board.ID == "" {
		_, err := r.db.Model(&model).Insert()
		if err != nil {
			return errors.Wrap(err, "insert board error")
		}
	} else {
		_, err := r.db.Model(&model).Update()
		if err != nil {
			return errors.Wrap(err, "update board error")
		}
	}

	board.ID = model.ID.String()
	board.CreatedAt = model.CreatedAt
	board.UpdatedAt = model.UpdatedAt
	board.DeletedAt = model.DeletedAt

	return nil
}
