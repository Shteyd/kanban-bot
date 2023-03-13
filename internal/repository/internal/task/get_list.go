package task

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *taskRepository) GetList(boardId string) ([]entity.Task, error) {
	var models model.TaskList

	err := r.db.Model(&models).Where("board_id = ?", uuid.MustParse(boardId)).Select()
	if err != nil {
		return nil, errors.Wrap(err, "select tasks list error")
	}

	return models.TaskListEntity(), nil
}
