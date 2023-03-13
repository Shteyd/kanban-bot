package task

import (
	"kanban-bot/internal/core/entity"
	"kanban-bot/internal/repository/internal/model"

	"github.com/pkg/errors"
)

func (r *taskRepository) Save(task *entity.Task) error {
	model := model.NewTask(*task)

	if model.ID.String() == "" {
		_, err := r.db.Model(&model).Insert()
		if err != nil {
			return errors.Wrap(err, "insert task error")
		}

		task.ID = model.ID.String()
	} else {
		_, err := r.db.Model(&model).WherePK().Update()
		if err != nil {
			return errors.Wrap(err, "update task error")
		}
	}

	return nil
}
