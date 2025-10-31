package repository

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/constants"
)

// FindByID finds a transaction by its ID
func (r *Repository) FindByID(ctx context.Context, id string) (*schemas.Transaction, error) {
	var transaction schemas.Transaction
	err := r.DB.WithContext(ctx).First(&transaction, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// FindAll retrieves all transactions
func (r *Repository) FindAll(ctx context.Context) ([]schemas.Transaction, error) {
	var transactions []schemas.Transaction
	err := r.DB.WithContext(ctx).Find(&transactions).Error
	return transactions, err
}

// FindByStatus retrieves all transactions with a specific status
func (r *Repository) FindByStatus(ctx context.Context, status schemas.TransactionStatus) ([]schemas.Transaction, error) {
	var transactions []schemas.Transaction
	err := r.DB.WithContext(ctx).
		Where("status = ?", status).
		Order("timestamp DESC").
		Find(&transactions).Error
	return transactions, err
}

// GetBalance calculates the balance (credits - debits) for successful transactions only
func (r *Repository) GetBalance(ctx context.Context) (int64, int64, error) {
	var credits int64
	var debits int64

	// Calculate total credits (type = CREDIT and status = SUCCESS)
	err := r.DB.WithContext(ctx).
		Model(&schemas.Transaction{}).
		Where("type = ? AND status = ?", schemas.TypeCredit, schemas.StatusSuccess).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&credits)
	if err != nil {
		return 0, 0, err
	}

	// Calculate total debits (type = DEBIT and status = SUCCESS)
	err = r.DB.WithContext(ctx).
		Model(&schemas.Transaction{}).
		Where("type = ? AND status = ?", schemas.TypeDebit, schemas.StatusSuccess).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&debits)
	if err != nil {
		return 0, 0, err
	}

	return credits - debits, credits, err
}

// GetIssuesWithFiltersAndSort retrieves transactions with filtering, sorting, and pagination
func (r *Repository) GetIssuesWithFiltersAndSort(
	ctx context.Context,
	page, pageSize int,
	filters schemas.TransactionFilters,
	sort schemas.TransactionSort,
) (*schemas.IssuesResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	var transactions []schemas.Transaction
	var total int64

	// Build query
	query := r.DB.WithContext(ctx).
		Where("status IN (?, ?)", schemas.StatusFailed, schemas.StatusPending)

	// Apply filters
	if filters.Status != "" {
		query = query.Where("status = ?", strings.ToUpper(filters.Status))
	}

	if filters.Type != "" {
		query = query.Where("type = ?", strings.ToUpper(filters.Type))
	}

	if filters.Amount > 0 {
		query = query.Where("amount = ?", filters.Amount)
	}

	if filters.SearchQuery != "" {
		searchQuery := "%" + filters.SearchQuery + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", searchQuery, searchQuery)
	}

	if filters.StartDate != "" && filters.EndDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", filters.StartDate, filters.EndDate)
	} else if filters.StartDate != "" {
		query = query.Where("DATE(created_at) >= ?", filters.StartDate)
	} else if filters.EndDate != "" {
		query = query.Where("DATE(created_at) <= ?", filters.EndDate)
	}

	// Get total count
	err := query.Model(&schemas.Transaction{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// Apply sorting only if both sort field and order are provided
	if sort.By != "" && sort.Order != "" {
		sortOrder := strings.ToUpper(sort.Order)
		query = query.Order(fmt.Sprintf("%s %s", sort.By, sortOrder))
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	err = query.
		Offset(offset).
		Limit(pageSize).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	// Convert to IssueTransaction format
	issues := make([]schemas.IssueTransaction, len(transactions))
	for i, t := range transactions {
		issues[i] = schemas.IssueTransaction{
			ID:          t.ID,
			Timestamp:   t.Timestamp,
			Name:        t.Name,
			Type:        string(t.Type),
			Amount:      t.Amount,
			Status:      string(t.Status),
			Description: t.Description,
			CreatedAt:   t.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// Build pagination links
	var nextLink *string
	var prevLink *string

	if page < totalPages {
		nextURL := fmt.Sprintf("?page=%d&page_size=%d", page+1, pageSize)
		nextLink = &nextURL
	}

	if page > 1 {
		prevURL := fmt.Sprintf("?page=%d&page_size=%d", page-1, pageSize)
		prevLink = &prevURL
	}

	// Build filter metadata
	filtersMeta := make(map[string]interface{})
	if filters.Status != "" {
		filtersMeta["status"] = filters.Status
	}
	if filters.Type != "" {
		filtersMeta["type"] = filters.Type
	}
	if filters.SearchQuery != "" {
		filtersMeta["search"] = filters.SearchQuery
	}
	if filters.Amount > 0 {
		filtersMeta["amount"] = filters.Amount
	}
	if filters.StartDate != "" {
		filtersMeta["start_date"] = filters.StartDate
	}
	if filters.EndDate != "" {
		filtersMeta["end_date"] = filters.EndDate
	}

	return &schemas.IssuesResponse{
		Message: constants.MsgIssuesRetrieved,
		Data:    issues,
		Meta: schemas.ResponseMeta{
			Pagination: schemas.PaginationMeta{
				Total:       int(total),
				Count:       len(issues),
				PerPage:     pageSize,
				CurrentPage: page,
				TotalPages:  totalPages,
				Links: schemas.PaginationLinks{
					Next: nextLink,
					Prev: prevLink,
				},
			},
			Filters: filtersMeta,
			Sort: &schemas.SortMeta{
				By:    sort.By,
				Order: sort.Order,
			},
		},
	}, nil
}

// GetIssues retrieves all non-successful transactions (FAILED + PENDING) with pagination
func (r *Repository) GetIssues(ctx context.Context, page int, pageSize int) (*schemas.IssuesResponse, error) {
	// Use default filters and sort
	filters := schemas.TransactionFilters{}
	sort := schemas.TransactionSort{
		By:    "timestamp",
		Order: "DESC",
	}
	return r.GetIssuesWithFiltersAndSort(ctx, page, pageSize, filters, sort)
}

// GetAllWithFiltersAndSort retrieves all transactions with filtering, sorting, and pagination
func (r *Repository) GetAllWithFiltersAndSort(
	ctx context.Context,
	page, pageSize int,
	filters schemas.TransactionFilters,
	sort schemas.TransactionSort,
) (*schemas.IssuesResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	var transactions []schemas.Transaction
	var total int64

	// Build query - fetch ALL transactions, not just issues
	query := r.DB.WithContext(ctx)

	// Apply filters
	if filters.Status != "" {
		query = query.Where("status = ?", strings.ToUpper(filters.Status))
	}

	if filters.Type != "" {
		query = query.Where("type = ?", strings.ToUpper(filters.Type))
	}

	if filters.Amount > 0 {
		query = query.Where("amount = ?", filters.Amount)
	}

	if filters.SearchQuery != "" {
		searchQuery := "%" + filters.SearchQuery + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", searchQuery, searchQuery)
	}

	if filters.StartDate != "" && filters.EndDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", filters.StartDate, filters.EndDate)
	} else if filters.StartDate != "" {
		query = query.Where("DATE(created_at) >= ?", filters.StartDate)
	} else if filters.EndDate != "" {
		query = query.Where("DATE(created_at) <= ?", filters.EndDate)
	}

	// Get total count
	err := query.Model(&schemas.Transaction{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// Apply sorting only if both sort field and order are provided
	if sort.By != "" && sort.Order != "" {
		sortOrder := strings.ToUpper(sort.Order)
		query = query.Order(fmt.Sprintf("%s %s", sort.By, sortOrder))
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	err = query.
		Offset(offset).
		Limit(pageSize).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	// Convert to IssueTransaction format
	issues := make([]schemas.IssueTransaction, len(transactions))
	for i, t := range transactions {
		issues[i] = schemas.IssueTransaction{
			ID:          t.ID,
			Timestamp:   t.Timestamp,
			Name:        t.Name,
			Type:        string(t.Type),
			Amount:      t.Amount,
			Status:      string(t.Status),
			Description: t.Description,
			CreatedAt:   t.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// Build pagination links
	var nextLink *string
	var prevLink *string

	if page < totalPages {
		nextURL := fmt.Sprintf("?page=%d&page_size=%d", page+1, pageSize)
		nextLink = &nextURL
	}

	if page > 1 {
		prevURL := fmt.Sprintf("?page=%d&page_size=%d", page-1, pageSize)
		prevLink = &prevURL
	}

	// Build filter metadata
	filtersMeta := make(map[string]interface{})
	if filters.Status != "" {
		filtersMeta["status"] = filters.Status
	}
	if filters.Type != "" {
		filtersMeta["type"] = filters.Type
	}
	if filters.SearchQuery != "" {
		filtersMeta["search"] = filters.SearchQuery
	}
	if filters.Amount > 0 {
		filtersMeta["amount"] = filters.Amount
	}
	if filters.StartDate != "" {
		filtersMeta["start_date"] = filters.StartDate
	}
	if filters.EndDate != "" {
		filtersMeta["end_date"] = filters.EndDate
	}

	return &schemas.IssuesResponse{
		Message: constants.MsgTransactionsRetrieved,
		Data:    issues,
		Meta: schemas.ResponseMeta{
			Pagination: schemas.PaginationMeta{
				Total:       int(total),
				Count:       len(issues),
				PerPage:     pageSize,
				CurrentPage: page,
				TotalPages:  totalPages,
				Links: schemas.PaginationLinks{
					Next: nextLink,
					Prev: prevLink,
				},
			},
			Filters: filtersMeta,
			Sort: &schemas.SortMeta{
				By:    sort.By,
				Order: sort.Order,
			},
		},
	}, nil
}

// CountByStatus counts transactions with a specific status
func (r *Repository) CountByStatus(ctx context.Context, status schemas.TransactionStatus) (int64, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&schemas.Transaction{}).
		Where("status = ?", status).
		Count(&count).Error
	return count, err
}

// Count returns the total number of transactions
func (r *Repository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&schemas.Transaction{}).
		Count(&count).Error
	return count, err
}
