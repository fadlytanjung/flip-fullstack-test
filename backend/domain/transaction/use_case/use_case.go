package use_case

import (
	"context"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/repository"
	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
)

// IUseCase defines the contract for transaction use case operations
type IUseCase interface {
	GetBalance(ctx context.Context) (*schemas.BalanceResponse, error)
	GetIssues(ctx context.Context, page int, pageSize int) (*schemas.IssuesResponse, error)
	GetIssuesWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error)
	GetAllWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error)
}

// UseCase implements IUseCase
type UseCase struct {
	Repository repository.IRepository
}

// NewUseCase creates a new transaction use case instance
func NewUseCase(repo repository.IRepository) IUseCase {
	return &UseCase{
		Repository: repo,
	}
}

// GetBalance calculates the balance from successful transactions
func (uc *UseCase) GetBalance(ctx context.Context) (*schemas.BalanceResponse, error) {
	balance, credits, err := uc.Repository.GetBalance(ctx)
	if err != nil {
		return nil, err
	}

	// Calculate debits from balance and credits
	debits := credits - balance

	return &schemas.BalanceResponse{
		Balance: balance,
		Credits: credits,
		Debits:  debits,
	}, nil
}

// GetIssues retrieves non-successful transactions
func (uc *UseCase) GetIssues(ctx context.Context, page int, pageSize int) (*schemas.IssuesResponse, error) {
	return uc.Repository.GetIssues(ctx, page, pageSize)
}

// GetIssuesWithFiltersAndSort retrieves issues with filtering and sorting
func (uc *UseCase) GetIssuesWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error) {
	return uc.Repository.GetIssuesWithFiltersAndSort(ctx, page, pageSize, filters, sort)
}

// GetAllWithFiltersAndSort retrieves all transactions with filtering and sorting
func (uc *UseCase) GetAllWithFiltersAndSort(ctx context.Context, page int, pageSize int, filters schemas.TransactionFilters, sort schemas.TransactionSort) (*schemas.IssuesResponse, error) {
	return uc.Repository.GetAllWithFiltersAndSort(ctx, page, pageSize, filters, sort)
}
