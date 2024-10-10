// Package repository provides the database operations for the todo endpoint.
package repository

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

// Todo is the repository for the todo endpoint.
type Todo interface {
	Create(t *model.Todo) error
	Delete(id int) error
	Update(t *model.Todo) error
	Find(id int) (*model.Todo, error)
	FindAll(queryParam *model.QueryParam) ([]*model.Todo, int64, error)
}

type todo struct {
	db *gorm.DB
}

// NewTodo returns a new instance of the todo repository.
func NewTodo(db *gorm.DB) Todo {
	return &todo{
		db: db,
	}
}

func (td *todo) Create(t *model.Todo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Update(t *model.Todo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Delete(id int) error {
	result := td.db.Where("id = ?", id).Delete(&model.Todo{})
	if result.RowsAffected == 0 {
		return model.ErrNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	log.Info("Deleted todo with id: ", id)
	return nil
}

func (td *todo) Find(id int) (*model.Todo, error) {
	var todo *model.Todo
	err := td.db.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (td *todo) FindAll(queryParam *model.QueryParam) ([]*model.Todo, int64, error) {
	var todos []*model.Todo
	var total int64

	query := td.db.Model(&model.Todo{})

	// Build the query for filtering
	if queryParam.Task != "" {
		query = query.Where("LOWER(task) LIKE ?", "%"+strings.ToLower(queryParam.Task)+"%")
	}
	if queryParam.Status != "" {
		query = query.Where("LOWER(status) LIKE ?", "%"+strings.ToLower(queryParam.Status)+"%")
	}

	// Count the total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Order and apply pagination
	query = query.Order("priority ASC")
	if queryParam.Page > 0 && queryParam.PerPage > 0 {
		query = query.Offset((queryParam.Page - 1) * queryParam.PerPage).Limit(queryParam.PerPage)
	}

	// Fetch the todos
	if err := query.Find(&todos).Error; err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}
