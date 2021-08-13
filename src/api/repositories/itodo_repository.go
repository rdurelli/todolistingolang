package repositories

import "github.com/rdurelli/todolist/src/api/models"

type Writer interface {
	Save(todo *models.Todo) (*models.Todo, error)
	Update(id string, todo *models.Todo) (*models.Todo, error)
}

type Reader interface {
	FindById(id uint) (*models.Todo, error)
	FindAll() (*[]*models.Todo, error)
	FindAllDone() (*[]*models.Todo, error)
}

type ITodoRepository interface {
	Writer
	Reader
}
