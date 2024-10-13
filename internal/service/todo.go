// Package service provides the business logic for the todo endpoint.
package service

import (
	"fmt"
	"time"

	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
)

// Todo is the service for the todo endpoint.
type Todo interface {
	Create(req *model.CreateRequest) (*model.Todo, error)
	Update(req model.UpdateRequest) (*model.Todo, error)
	Delete(id int) (*model.Todo, error)
	Find(id int) (*model.Todo, error)
	FindAll(queryParam *model.QueryParam) ([]*model.Todo, int64, error)
	FindStatusList(queryParam *model.QueryParam) ([]*model.Todo, int64, error)
	FindAllTasksWithoutPagination() ([]*model.Todo, error)
}

type todo struct {
	todoRepository repository.Todo
}

// NewTodo creates a new Todo service.
func NewTodo(r repository.Todo) Todo {
	return &todo{r}
}

func (t *todo) Create(req *model.CreateRequest) (*model.Todo, error) {
	todo := model.NewTodo(req)
	value, ok := model.PriorityMap[string(todo.Status)]
	fmt.Println("task status: ", todo.Status)
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

func (t *todo) Update(req model.UpdateRequest) (*model.Todo, error) {
	currentTodo, err := t.Find(req.ID)
	if err != nil {
		return nil, err
	}

	// 空文字列の場合、現在の値を使用
	if req.TaskTitle != "" {
		currentTodo.Task = req.TaskTitle
	}
	if req.Description != "" {
		currentTodo.Description = req.Description
	}
	if req.Status != "" {
		currentTodo.Status = req.Status
		value, ok := model.PriorityMap[string(currentTodo.Status)]
		if !ok {
			return nil, model.ErrStatusNotFound
		}
		currentTodo.Priority = value
	}
	if req.PriorityTask != "" {
		currentTodo.PriorityTask = req.PriorityTask
	}
	currentTodo.UpdatedAt = time.Now().UTC()

	if err := t.todoRepository.Update(currentTodo); err != nil {
		return nil, err
	}
	return currentTodo, nil
}

func (t *todo) Delete(id int) (*model.Todo, error) {
	currentTodo, err := t.Find(id)
	if err != nil {
		return nil, err
	}
	if err := t.todoRepository.Delete(id); err != nil {
		return nil, err
	}
	return currentTodo, nil
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

func (t *todo) FindStatusList(queryParam *model.QueryParam) ([]*model.Todo, int64, error) {
	todos, total, err := t.todoRepository.FindStatusList(queryParam)
	if err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}

func (t *todo) FindAllTasksWithoutPagination() ([]*model.Todo, error) {
	todos, err := t.todoRepository.FindAllTasksWithoutPagination()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
