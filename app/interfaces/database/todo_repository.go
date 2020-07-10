package database

import "github.com/horizon67/go-clean-arch-todo-example/app/domain"

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(id int) (todo domain.Todo, err error) {
	if err = repo.Find(&todo, id).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {
	if err = repo.Find(&todos).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) Store(t domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Create(&t).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) Update(t domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Save(&t).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) DeleteById(todo domain.Todo) (err error) {
	if err = repo.Delete(&todo).Error; err != nil {
		return
	}
	return
}
