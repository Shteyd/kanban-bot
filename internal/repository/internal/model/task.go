package model

import (
	"kanban-bot/internal/core/entity"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	BoardID     string    `json:"board_id"`
	Position    int       `json:"position"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TaskList []Task

func NewTask(entity entity.Task) Task {
	return Task{
		ID:          uuid.MustParse(entity.ID),
		BoardID:     entity.BoardID,
		Position:    entity.Position,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}

func (m *Task) TaskEntity() entity.Task {
	return entity.Task{
		ID:          m.ID.String(),
		BoardID:     m.BoardID,
		Position:    m.Position,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
}

func (l *TaskList) TaskListEntity() []entity.Task {
	var entities []entity.Task

	for _, v := range *l {
		entities = append(entities, v.TaskEntity())
	}

	return entities
}
