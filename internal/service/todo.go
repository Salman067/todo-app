// Package service provides the business logic for the todo endpoint.
package service

import (
	"time"

	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
)

// Todo is the service for the todo endpoint.
type Todo interface {
	Create(task string) (*model.Todo, error)
	Update(id int, task string, status model.Status) (*model.Todo, error)
	Delete(id int) error
	Find(id int) (*model.Todo, error)
	FindAll(queryParam *model.QueryParam) ([]*model.Todo, int64, error)
}

type todo struct {
	todoRepository repository.Todo
}

// NewTodo creates a new Todo service.
func NewTodo(r repository.Todo) Todo {
	return &todo{r}
}

func (t *todo) Create(task string) (*model.Todo, error) {
	todo := model.NewTodo(task)
	value, ok := model.PriorityMap[string(todo.Status)]
	if !ok {
		return nil, model.ErrStatusNotFound
	}
	todo.Priority = value
	todo.CreatedAt = time.Now().UTC()
	if err := t.todoRepository.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) Update(id int, task string, status model.Status) (*model.Todo, error) {
	currentTodo, err := t.Find(id)
	if err != nil {
		return nil, err
	}

	// 空文字列の場合、現在の値を使用
	if task != "" {
		currentTodo.Task = task
	}
	if status != "" {
		currentTodo.Status = status
		value, ok := model.PriorityMap[string(currentTodo.Status)]
		if !ok {
			return nil, model.ErrStatusNotFound
		}
		currentTodo.Priority = value
	}
	currentTodo.UpdatedAt = time.Now().UTC()

	if err := t.todoRepository.Update(currentTodo); err != nil {
		return nil, err
	}
	return currentTodo, nil
}

func (t *todo) Delete(id int) error {
	if err := t.todoRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *todo) Find(id int) (*model.Todo, error) {
	todo, err := t.todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) FindAll(queryParam *model.QueryParam) ([]*model.Todo, int64, error) {
	todos, total, err := t.todoRepository.FindAll(queryParam)
	if err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}
