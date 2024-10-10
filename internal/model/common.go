package model

import "fmt"

// ErrNotFound is the error for not found.
var ErrNotFound = fmt.Errorf("not found")
var ErrStatusNotFound = fmt.Errorf("status not found")

type QueryParam struct {
	Task    string `query:"task"`
	Status  string `query:"status"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}
