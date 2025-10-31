package repository

import (
	"context"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"gorm.io/gorm"
)

// IRepository defines the contract for transaction repository operations
type IRepository interface {
	// Commands
	Create(ctx context.Context, transaction *schemas.Transaction) error
	CreateBatch(ctx context.Context, transactions []schemas.Transaction) error
	DeleteAll(ctx context.Context) error

	// Queries
	FindByID(ctx context.Context, id string) (*schemas.Transaction, error)
	FindAll(ctx context.Context) ([]schemas.Transaction, error)
	FindByStatus(ctx context.Context, status schemas.TransactionStatus) ([]schemas.Transaction, error)
	GetBalance(ctx context.Context) (int64, int64, error)
	GetIssues(ctx context.Context, page int, pageSize int) (*schemas.IssuesResponse, error)
	GetIssuesWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error)
	GetAllWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error)
	CountByStatus(ctx context.Context, status schemas.TransactionStatus) (int64, error)
	Count(ctx context.Context) (int64, error)
}

// Repository implements IRepository
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new transaction repository instance
func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		DB: db,
	}
}
