package board

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *boardRepository) Get(id string) (entity.Board, error) {
	model := model.NewBoard(entity.Board{ID: id})

	if err := r.db.Model(&model).WherePK().Select(); err != nil {
		return entity.Board{}, errors.Wrap(err, "select board error")
	}

	return model.BoardEntity(), nil
}
