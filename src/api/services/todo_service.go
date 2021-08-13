package services

import (
	"github.com/rdurelli/todolist/src/api/lib"
	"github.com/rdurelli/todolist/src/api/models"
	"github.com/rdurelli/todolist/src/api/repositories"
)

type TodoService struct {
	logger lib.Logger
	repository repositories.ITodoRepository
}

func NewTodoService(logger lib.Logger, repository repositories.ITodoRepository) TodoService {
	logger.Info("Chamou o TODO SERVICE")
	return TodoService{
		logger:     logger,
		repository: repository,
	}
}

func (tS TodoService) GetOneTodo (id uint) (models.Todo, error) {
	return models.Todo{}, nil
}
