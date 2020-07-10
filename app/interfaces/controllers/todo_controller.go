package controllers

import (
	"github.com/horizon67/go-clean-arch-todo-example/app/domain"
	"github.com/horizon67/go-clean-arch-todo-example/app/interfaces/database"
	"github.com/horizon67/go-clean-arch-todo-example/app/usecase"
	"strconv"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Index(c Context) (err error) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todos)
	return
}

func (controller *TodoController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todo)
	return
}

func (controller *TodoController) Create(c Context) (err error) {
	t := domain.Todo{}
	c.Bind(&t)
	todo, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) Save(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	t, err := controller.Interactor.TodoById(id)
	c.Bind(&t)
	todo, err := controller.Interactor.Update(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := domain.Todo{
		ID: id,
	}
	err = controller.Interactor.DeleteById(todo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todo)
	return
}
