package repository

import (
	"kanban-bot/internal/core/interfaces"
	"kanban-bot/internal/repository/internal/board"
	"kanban-bot/internal/repository/internal/customer"
	"kanban-bot/internal/repository/internal/task"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"
)

type repositoryManager struct {
	DB *pg.DB
	Tx *pg.Tx
}

func NewRepositoryManager(db *pg.DB, tx *pg.Tx) interfaces.RepositoryManager {
	return &repositoryManager{
		DB: db,
		Tx: tx,
	}
}

func (rm *repositoryManager) BeginTx() (interfaces.RepositoryManager, error) {
	tx, err := rm.DB.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction error")
	}
	return NewRepositoryManager(rm.DB, tx), nil
}

func (rm *repositoryManager) CommitTx() error {
	defer rm.clearTransaction()
	return rm.Tx.Commit()
}

func (rm *repositoryManager) RollbackTx() error {
	defer rm.clearTransaction()
	return rm.Tx.Rollback()
}

func (rm *repositoryManager) Transaction(
	callbaсk func(interfaces.RepositoryManager) error,
) error {
	tx, err := rm.BeginTx()
	if err != nil {
		return errors.Wrap(err, "begin transaction error")
	}

	if err := callbaсk(tx); err != nil {
		rollbackErr := tx.RollbackTx()
		if rollbackErr != nil {
			return errors.Wrap(err, "rollback err")
		}
		return errors.Wrap(err, "callback execution err")
	}

	return tx.CommitTx()
}

func (rm *repositoryManager) clearTransaction() {
	rm.Tx = nil
}

func (rm *repositoryManager) getConnect() orm.DB {
	if rm.Tx == nil {
		return rm.DB
	}

	return rm.Tx
}

func (rm *repositoryManager) GetBoardRepository() interfaces.BoardRepository {
	return board.NewBoardRepository(rm.getConnect())
}

func (rm *repositoryManager) GetCustomerRepository() interfaces.CustomerRepository {
	return customer.NewCustomerRepository(rm.getConnect())
}

func (rm *repositoryManager) GetTaskRepository() interfaces.TaskRepository {
	return task.NewTaskRepository(rm.getConnect())
}
