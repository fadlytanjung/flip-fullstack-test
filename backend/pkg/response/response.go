package response

import "strconv"

// Status represents HTTP status codes
type Status int

const (
	StatusOK            Status = 200
	StatusCreated       Status = 201
	StatusBadRequest    Status = 400
	StatusUnauthorized  Status = 401
	StatusForbidden     Status = 403
	StatusNotFound      Status = 404
	StatusConflict      Status = 409
	StatusInternalError Status = 500
)

// Meta contains pagination and filter metadata
type Meta struct {
	Pagination *PaginationMeta         `json:"pagination,omitempty"`
	Filters    map[string]interface{}  `json:"filters,omitempty"`
	Sort       *SortMeta               `json:"sort,omitempty"`
	Message    string                  `json:"message,omitempty"`
}

// PaginationMeta represents pagination information
type PaginationMeta struct {
	Total       int               `json:"total"`
	Count       int               `json:"count"`
	PerPage     int               `json:"per_page"`
	CurrentPage int               `json:"current_page"`
	TotalPages  int               `json:"total_pages"`
	Links       PaginationLinks   `json:"links"`
}

// PaginationLinks represents pagination navigation links
type PaginationLinks struct {
	Next *string `json:"next"`
	Prev *string `json:"prev"`
}

// SortMeta represents sort information
type SortMeta struct {
	By    string `json:"by"`
	Order string `json:"order"`
}

// Success represents a successful API response
type Success struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Meta   *Meta       `json:"meta,omitempty"`
}

// Error represents an error API response
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// NewSuccess creates a new success response
func NewSuccess(data interface{}, meta *Meta) *Success {
	return &Success{
		Status: int(StatusOK),
		Data:   data,
		Meta:   meta,
	}
}

// NewSuccessWithStatus creates a new success response with custom status
func NewSuccessWithStatus(status Status, data interface{}, meta *Meta) *Success {
	return &Success{
		Status: int(status),
		Data:   data,
		Meta:   meta,
	}
}

// NewError creates a new error response
func NewError(status Status, message string, err string) *Error {
	return &Error{
		Status:  int(status),
		Message: message,
		Error:   err,
	}
}

// NewBadRequest creates a bad request error response
func NewBadRequest(message string, err string) *Error {
	return NewError(StatusBadRequest, message, err)
}

// NewInternalError creates an internal server error response
func NewInternalError(message string, err string) *Error {
	return NewError(StatusInternalError, message, err)
}

// NewNotFound creates a not found error response
func NewNotFound(message string) *Error {
	return NewError(StatusNotFound, message, "")
}

// NewPaginationMeta creates pagination metadata
func NewPaginationMeta(total, count, perPage, currentPage int) *PaginationMeta {
	totalPages := (total + perPage - 1) / perPage
	if totalPages < 1 {
		totalPages = 1
	}

	meta := &PaginationMeta{
		Total:       total,
		Count:       count,
		PerPage:     perPage,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}

	// Set navigation links
	if currentPage > 1 {
		prevPage := "page=" + strconv.Itoa(currentPage-1)
		meta.Links.Prev = &prevPage
	}

	if currentPage < totalPages {
		nextPage := "page=" + strconv.Itoa(currentPage+1)
		meta.Links.Next = &nextPage
	}

	return meta
}

// NewMeta creates response metadata with message
func NewMeta(message string, pagination *PaginationMeta, filters map[string]interface{}, sort *SortMeta) *Meta {
	return &Meta{
		Message:    message,
		Pagination: pagination,
		Filters:    filters,
		Sort:       sort,
	}
}
