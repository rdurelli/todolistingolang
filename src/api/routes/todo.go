package routes

import (
	"github.com/rdurelli/todolist/src/api/controllers"
	"github.com/rdurelli/todolist/src/api/lib"
)

type TodoRoutes struct {
	logger lib.Logger
	handler lib.RequestHandlerGin
	todoController controllers.TodoController
}

func (tR TodoRoutes) Setup() {
	tR.logger.Info("Setting up TodoRoutes")
	api := tR.handler.Gin.Group("/api")
	{
		api.GET("/todos", tR.todoController.GetOneTodo)
	}
}

func NewTodoRoutes(logger lib.Logger, handler lib.RequestHandlerGin, todoController controllers.TodoController) TodoRoutes{
	return TodoRoutes{
		logger:         logger,
		handler:        handler,
		todoController: todoController,
	}
}
