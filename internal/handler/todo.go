package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/errors"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/service"
)

// TodoHandler is the request handler for the todo endpoint.
type TodoHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Find(c echo.Context) error
	FindAll(c echo.Context) error
	FindStatusList(c echo.Context) error
	FindAllTasksWithoutPagination(c echo.Context) error
	handleError(c echo.Context, err error) error
	fetchTodos(c echo.Context, queryFunc func(*model.QueryParam) ([]*model.Todo, int64, error)) error
}

type todoHandler struct {
	Handler
	service service.Todo
}

// NewTodo returns a new instance of the todo handler.
func NewTodo(s service.Todo) TodoHandler {
	return &todoHandler{service: s}
}

// @Summary	Create a new todo
// @Tags		todos
// @Accept		json
// @Produce	json
// @Param		request	body		CreateRequest	true	"json"
// @Success	201		{object}	ResponseError{data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos [post]
func (t *todoHandler) Create(c echo.Context) error {
	var req model.CreateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Create(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusCreated, ResponseData{Data: todo})
}

// @Summary	Update a todo
// @Tags		todos
// @Accept		json
// @Produce	json
// @Param		body	body		UpdateRequestBody	true	"body"
// @Param		path	path		UpdateRequestPath	false	"path"
// @Success	201		{object}	ResponseData{Data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos/:id [put]
func (t *todoHandler) Update(c echo.Context) error {
	var req model.UpdateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Update(req)
	if err != nil {
		if err == model.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				ResponseError{Errors: []Error{{Code: errors.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusOK, ResponseData{Data: todo})
}

// DeleteRequest is the request parameter for deleting a todo
type DeleteRequest struct {
	ID int `param:"id" validate:"required"`
}

// @Summary	Delete a todo
// @Tags		todos
// @Param		path	path	DeleteRequest	false	"path"
// @Success	204
// @Failure	400	{object}	ResponseError
// @Failure	404	{object}	ResponseError
// @Failure	500	{object}	ResponseError
// @Router		/todos/:id [delete]
func (t *todoHandler) Delete(c echo.Context) error {
	var req DeleteRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Delete(req.ID)
	if err != nil {
		return t.handleError(c, err)
	}
	return c.JSON(http.StatusOK, ResponseData{Data: todo})
}

// FindRequest is the request parameter for finding a todo
type FindRequest struct {
	ID int `param:"id" validate:"required"`
}

// @Summary	Find a todo
// @Tags		todos
// @Param		path	path		FindRequest	false	"path"
// @Success	200		{object}	ResponseData{Data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	404		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos/:id [get]
func (t *todoHandler) Find(c echo.Context) error {
	var req FindRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	res, err := t.service.Find(req.ID)
	if err != nil {
		return t.handleError(c, err)
	}
	return c.JSON(http.StatusOK, ResponseData{Data: res})
}

// @Summary	Find all todos
// @Tags		todos
// @Success	200	{object}	ResponseData{Data=[]model.Todo}
// @Failure	500	{object}	ResponseError
// @Router		/todos [get]
func (t *todoHandler) FindAll(c echo.Context) error {
	return t.fetchTodos(c, func(queryParam *model.QueryParam) ([]*model.Todo, int64, error) {
		todos, total, err := t.service.FindAll(queryParam)
		if err != nil {
			return nil, 0, err
		}

		return todos, total, nil
	})
}

func (t *todoHandler) FindStatusList(c echo.Context) error {
	return t.fetchTodos(c, func(queryParam *model.QueryParam) ([]*model.Todo, int64, error) {
		todos, total, err := t.service.FindStatusList(queryParam)
		if err != nil {
			return nil, 0, err
		}
		return todos, total, nil
	})
}

func (t *todoHandler) FindAllTasksWithoutPagination(c echo.Context) error {
	res, err := t.service.FindAllTasksWithoutPagination()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, ResponseData{Data: res})
}

func (t *todoHandler) handleError(c echo.Context, err error) error {
	if err == model.ErrNotFound {
		return c.JSON(http.StatusNotFound,
			ResponseError{Errors: []Error{{Code: errors.CodeNotFound, Message: "todo not found"}}})
	}
	return c.JSON(http.StatusInternalServerError,
		ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
}

func (t *todoHandler) fetchTodos(c echo.Context, queryFunc func(*model.QueryParam) ([]*model.Todo, int64, error)) error {
	var queryParam model.QueryParam
	if err := c.Bind(&queryParam); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	// Default pagination
	if queryParam.Page <= 0 {
		queryParam.Page = 1
	}
	if queryParam.PerPage <= 0 {
		queryParam.PerPage = 10
	}

	res, total, err := queryFunc(&queryParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     res,
		"total":    total,
		"page":     queryParam.Page,
		"per_page": queryParam.PerPage,
	})
}
