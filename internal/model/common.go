package model

import "fmt"

// ErrNotFound is the error for not found.
var (
	ErrNotFound       = fmt.Errorf("not found")
	ErrStatusNotFound = fmt.Errorf("status not found")
)

// QueryParam represents the query parameters for pagination and filtering
// in API requests. It includes options for page number, items per page,
// and a keyword for filtering.
type QueryParam struct {
	Page    int    `query:"page"`     // Page is the current page number for pagination.
	PerPage int    `query:"per_page"` // PerPage defines the number of items per page.
	KeyWord string `query:"key_word"` // KeyWord is an optional search term for filtering.
	Status  string `query:"status"`   // Status is an optional status filter.
}
