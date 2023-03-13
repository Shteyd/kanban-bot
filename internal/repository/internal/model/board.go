package model

import (
	"kanban-bot/internal/core/entity"
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID          uuid.UUID `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Article     string    `json:"article"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type BoardList []Board

func NewBoard(entity entity.Board) Board {
	return Board{
		ID:          uuid.MustParse(entity.ID),
		CustomerID:  entity.CustomerID,
		Article:     entity.Article,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}

func (m *Board) BoardEntity() entity.Board {
	return entity.Board{
		ID:          m.ID.String(),
		CustomerID:  m.CustomerID,
		Article:     m.Article,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
}

func (l *BoardList) BoardListEntity() []entity.Board {
	var entities []entity.Board

	for _, v := range *l {
		entities = append(entities, v.BoardEntity())
	}

	return entities
}
