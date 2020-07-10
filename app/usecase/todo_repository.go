package usecase

import "github.com/horizon67/go-clean-arch-todo-example/app/domain"

type TodoRepository interface {
	FindById(id int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
	Store(domain.Todo) (domain.Todo, error)
	Update(domain.Todo) (domain.Todo, error)
	DeleteById(domain.Todo) error
}
