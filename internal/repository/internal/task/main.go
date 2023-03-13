package task

import (
	"kanban-bot/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type taskRepository struct {
	db orm.DB
}

func NewTaskRepository(db orm.DB) interfaces.TaskRepository {
	return &taskRepository{db: db}
}
