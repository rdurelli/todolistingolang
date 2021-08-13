package repositories

import (
	"github.com/rdurelli/todolist/src/api/lib"
	"github.com/rdurelli/todolist/src/api/models"
)

type todoRepository struct {
	logger lib.Logger
}

func NewTodoRepository(logger lib.Logger) ITodoRepository {
	logger.Info("Chamou o TODO REPOSITORY")
	return &todoRepository{
		logger: logger,
	}
}

func (tR todoRepository) Save(todo *models.Todo) (*models.Todo, error) {
	panic("implement me")
}

func (tR todoRepository) Update(id string, todo *models.Todo) (*models.Todo, error) {
	panic("implement me")
}

func (tR todoRepository) FindById(id uint) (*models.Todo, error) {
	panic("implement me")
}

func (tR todoRepository) FindAll() (*[]*models.Todo, error) {
	panic("implement me")
}

func (tR todoRepository) FindAllDone() (*[]*models.Todo, error) {
	panic("implement me")
}


