package task

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *taskRepository) Get(id string) (entity.Task, error) {
	model := model.NewTask(entity.Task{ID: id})

	err := r.db.Model(&model).WherePK().Select()
	if err != nil {
		return entity.Task{}, errors.Wrap(err, "select task error")
	}

	return model.TaskEntity(), nil
}
