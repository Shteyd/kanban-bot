package interfaces

import "kanban-bot/internal/core/entity"

type (
	RepositoryManager interface {
		BeginTx() (RepositoryManager, error)
		RollbackTx() error
		CommitTx() error
		Transaction(func(RepositoryManager) error) error

		GetBoardRepository() BoardRepository
		GetCustomerRepository() CustomerRepository
		GetTaskRepository() TaskRepository
	}

	BoardRepository interface {
		Save(board *entity.Board) error
		Get(id string) (entity.Board, error)
		GetList(customerId string) ([]entity.Board, error)
	}

	CustomerRepository interface {
		Save(customer *entity.Customer) error
		Get(id string) (entity.Customer, error)
	}

	TaskRepository interface {
		Save(task *entity.Task) error
		Get(id string) (entity.Task, error)
		GetList(boardId string) ([]entity.Task, error)
	}
)
