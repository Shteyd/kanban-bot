package board

import (
	"kanban-bot/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type boardRepository struct {
	db orm.DB
}

func NewBoardRepository(db orm.DB) interfaces.BoardRepository {
	return &boardRepository{db: db}
}
