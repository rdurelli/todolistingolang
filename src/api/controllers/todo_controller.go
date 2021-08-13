package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rdurelli/todolist/src/api/lib"
	"github.com/rdurelli/todolist/src/api/services"
)

type TodoController struct {
	logger lib.Logger
	service services.TodoService
}

func NewTodoController(logger lib.Logger, services services.TodoService) TodoController {
	return TodoController{
		logger: logger,
		service: services,
	}
}


func (pC TodoController) GetOneTodo(c *gin.Context) {
	pC.logger.Info("Eu passei por aqui sim....")
	todo, err:=pC.service.GetOneTodo(1)
	if err != nil {
		panic(err)
	}
	pC.logger.Info("Eu passei por aqui sim....", todo)
	c.JSON(200, "Oiii")
}
