package schemas

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string
type TransactionStatus string

const (
	TypeCredit TransactionType = "CREDIT"
	TypeDebit  TransactionType = "DEBIT"
)

const (
	StatusSuccess TransactionStatus = "SUCCESS"
	StatusFailed  TransactionStatus = "FAILED"
	StatusPending TransactionStatus = "PENDING"
)

// Transaction represents a bank transaction
type Transaction struct {
	ID          string                `gorm:"primaryKey;type:text" json:"id"`
	Timestamp   int64                 `gorm:"index" json:"timestamp"`
	Name        string                `json:"name"`
	Type        TransactionType       `gorm:"type:text" json:"type"`
	Amount      int64                 `json:"amount"`
	Status      TransactionStatus     `gorm:"type:text;index" json:"status"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   gorm.DeletedAt        `gorm:"index" json:"-"`
}

// TableName specifies the table name for Transaction
func (Transaction) TableName() string {
	return "transactions"
}

// UploadRequest represents the upload CSV request
type UploadRequest struct {
	// File will be parsed from multipart form
}

// UploadResponse represents the response after upload
type UploadResponse struct {
	Message        string `json:"message"`
	TotalRecords   int    `json:"total_records"`
	SuccessRecords int    `json:"success_records"`
	FailedRecords  int    `json:"failed_records"`
	PendingRecords int    `json:"pending_records"`
}

// BalanceResponse represents the balance calculation response
type BalanceResponse struct {
	Balance int64 `json:"balance"`
	Credits int64 `json:"credits"`
	Debits  int64 `json:"debits"`
}

// IssueTransaction represents a non-successful transaction for issues endpoint
type IssueTransaction struct {
	ID          string `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// PaginationLinks represents pagination navigation links
type PaginationLinks struct {
	Next *string `json:"next"`
	Prev *string `json:"prev"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	Total      int               `json:"total"`
	Count      int               `json:"count"`
	PerPage    int               `json:"per_page"`
	CurrentPage int              `json:"current_page"`
	TotalPages int               `json:"total_pages"`
	Links      PaginationLinks   `json:"links"`
}

// ResponseMeta contains pagination and filter metadata
type ResponseMeta struct {
	Pagination PaginationMeta          `json:"pagination"`
	Filters    map[string]interface{}  `json:"filters,omitempty"`
	Sort       *SortMeta               `json:"sort,omitempty"`
}

// SortMeta represents sort information
type SortMeta struct {
	By    string `json:"by"`
	Order string `json:"order"`
}

// IssuesResponse represents the issues list response matching frontend contract
type IssuesResponse struct {
	Message string                 `json:"message"`
	Data    []IssueTransaction     `json:"data"`
	Meta    ResponseMeta           `json:"meta"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse represents a success response wrapper
type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta,omitempty"`
}

// TransactionFilters represents filtering options
type TransactionFilters struct {
	Status      string
	Type        string
	SearchQuery string
	Amount      int64
	StartDate   string
	EndDate     string
}

// TransactionSort represents sorting options
type TransactionSort struct {
	By    string // timestamp, amount, name, status, type, description, created_at
	Order string // ASC or DESC
}
