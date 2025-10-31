package repository

import (
	"context"
	"testing"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&schemas.Transaction{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	return db
}

// TestGetBalance tests balance calculation from successful transactions
func TestGetBalance(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	// Insert test data
	transactions := []schemas.Transaction{
		{
			ID:        "1",
			Timestamp: 1000,
			Name:      "Test Credit",
			Type:      schemas.TypeCredit,
			Amount:    1000000,
			Status:    schemas.StatusSuccess,
		},
		{
			ID:        "2",
			Timestamp: 2000,
			Name:      "Test Debit",
			Type:      schemas.TypeDebit,
			Amount:    300000,
			Status:    schemas.StatusSuccess,
		},
		{
			ID:        "3",
			Timestamp: 3000,
			Name:      "Failed Credit",
			Type:      schemas.TypeCredit,
			Amount:    500000,
			Status:    schemas.StatusFailed,
		},
		{
			ID:        "4",
			Timestamp: 4000,
			Name:      "Pending Debit",
			Type:      schemas.TypeDebit,
			Amount:    200000,
			Status:    schemas.StatusPending,
		},
	}

	if err := db.CreateInBatches(transactions, 100).Error; err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Test balance calculation
	balance, credits, err := repo.GetBalance(ctx)
	if err != nil {
		t.Fatalf("GetBalance failed: %v", err)
	}

	// Expected: balance = 1000000 (credit) - 300000 (debit) = 700000
	expectedBalance := int64(700000)
	expectedCredits := int64(1000000)

	if balance != expectedBalance {
		t.Errorf("Expected balance %d, got %d", expectedBalance, balance)
	}

	if credits != expectedCredits {
		t.Errorf("Expected credits %d, got %d", expectedCredits, credits)
	}
}

// TestGetIssues tests filtering of failed and pending transactions
func TestGetIssues(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	// Insert test data
	transactions := []schemas.Transaction{
		{
			ID:        "1",
			Timestamp: 1000,
			Name:      "Success Transaction",
			Type:      schemas.TypeCredit,
			Amount:    1000000,
			Status:    schemas.StatusSuccess,
		},
		{
			ID:        "2",
			Timestamp: 2000,
			Name:      "Failed Transaction",
			Type:      schemas.TypeDebit,
			Amount:    300000,
			Status:    schemas.StatusFailed,
		},
		{
			ID:        "3",
			Timestamp: 3000,
			Name:      "Pending Transaction",
			Type:      schemas.TypeDebit,
			Amount:    200000,
			Status:    schemas.StatusPending,
		},
	}

	if err := db.CreateInBatches(transactions, 100).Error; err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Test get issues
	response, err := repo.GetIssues(ctx, 1, 10)
	if err != nil {
		t.Fatalf("GetIssues failed: %v", err)
	}

	// Should return 2 issues (failed + pending)
	if response.Meta.Pagination.Total != 2 {
		t.Errorf("Expected 2 total issues, got %d", response.Meta.Pagination.Total)
	}

	if len(response.Data) != 2 {
		t.Errorf("Expected 2 issue transactions, got %d", len(response.Data))
	}

	// Verify success transaction is not included
	for _, issue := range response.Data {
		if issue.Status == "SUCCESS" {
			t.Error("Success transaction should not be in issues")
		}
	}
}

// TestGetIssuesWithFilters tests filtering with status and type filters
func TestGetIssuesWithFilters(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	// Insert test data
	transactions := []schemas.Transaction{
		{
			ID:        "1",
			Timestamp: 1000,
			Name:      "Failed Credit",
			Type:      schemas.TypeCredit,
			Amount:    1000000,
			Status:    schemas.StatusFailed,
		},
		{
			ID:        "2",
			Timestamp: 2000,
			Name:      "Failed Debit",
			Type:      schemas.TypeDebit,
			Amount:    300000,
			Status:    schemas.StatusFailed,
		},
		{
			ID:        "3",
			Timestamp: 3000,
			Name:      "Pending Debit",
			Type:      schemas.TypeDebit,
			Amount:    200000,
			Status:    schemas.StatusPending,
		},
	}

	if err := db.CreateInBatches(transactions, 100).Error; err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Test with status filter
	filters := schemas.TransactionFilters{
		Status: "FAILED",
	}
	sort := schemas.TransactionSort{
		By:    "timestamp",
		Order: "DESC",
	}

	response, err := repo.GetIssuesWithFiltersAndSort(ctx, 1, 10, filters, sort)
	if err != nil {
		t.Fatalf("GetIssuesWithFiltersAndSort failed: %v", err)
	}

	if response.Meta.Pagination.Total != 2 {
		t.Errorf("Expected 2 failed transactions, got %d", response.Meta.Pagination.Total)
	}

	// Test with type filter
	filters.Status = ""
	filters.Type = "DEBIT"

	response, err = repo.GetIssuesWithFiltersAndSort(ctx, 1, 10, filters, sort)
	if err != nil {
		t.Fatalf("GetIssuesWithFiltersAndSort failed: %v", err)
	}

	if response.Meta.Pagination.Total != 2 {
		t.Errorf("Expected 2 debit issues, got %d", response.Meta.Pagination.Total)
	}
}

// TestCountByStatus tests counting transactions by status
func TestCountByStatus(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	// Insert test data
	transactions := []schemas.Transaction{
		{ID: "1", Status: schemas.StatusSuccess},
		{ID: "2", Status: schemas.StatusSuccess},
		{ID: "3", Status: schemas.StatusFailed},
		{ID: "4", Status: schemas.StatusPending},
	}

	if err := db.CreateInBatches(transactions, 100).Error; err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Test count by status
	successCount, err := repo.CountByStatus(ctx, schemas.StatusSuccess)
	if err != nil {
		t.Fatalf("CountByStatus failed: %v", err)
	}

	if successCount != 2 {
		t.Errorf("Expected 2 success transactions, got %d", successCount)
	}

	failedCount, err := repo.CountByStatus(ctx, schemas.StatusFailed)
	if err != nil {
		t.Fatalf("CountByStatus failed: %v", err)
	}

	if failedCount != 1 {
		t.Errorf("Expected 1 failed transaction, got %d", failedCount)
	}
}

// TestPagination tests pagination logic
func TestPagination(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	// Insert 15 test transactions
	for i := 1; i <= 15; i++ {
		status := schemas.StatusFailed
		if i%3 == 0 {
			status = schemas.StatusPending
		}
		transaction := schemas.Transaction{
			ID:        string(rune(i)),
			Timestamp: int64(i * 1000),
			Status:    status,
		}
		if err := db.Create(&transaction).Error; err != nil {
			t.Fatalf("failed to insert test data: %v", err)
		}
	}

	// Test first page
	response, err := repo.GetIssues(ctx, 1, 5)
	if err != nil {
		t.Fatalf("GetIssues page 1 failed: %v", err)
	}

	if len(response.Data) != 5 {
		t.Errorf("Expected 5 items on page 1, got %d", len(response.Data))
	}

	if response.Meta.Pagination.CurrentPage != 1 {
		t.Errorf("Expected current page 1, got %d", response.Meta.Pagination.CurrentPage)
	}

	if response.Meta.Pagination.TotalPages != 3 {
		t.Errorf("Expected 3 total pages, got %d", response.Meta.Pagination.TotalPages)
	}

	// Verify next link exists
	if response.Meta.Pagination.Links.Next == nil {
		t.Error("Expected next link for page 1")
	}

	// Test second page
	response, err = repo.GetIssues(ctx, 2, 5)
	if err != nil {
		t.Fatalf("GetIssues page 2 failed: %v", err)
	}

	if len(response.Data) != 5 {
		t.Errorf("Expected 5 items on page 2, got %d", len(response.Data))
	}

	if response.Meta.Pagination.CurrentPage != 2 {
		t.Errorf("Expected current page 2, got %d", response.Meta.Pagination.CurrentPage)
	}

	// Verify both prev and next links exist
	if response.Meta.Pagination.Links.Prev == nil {
		t.Error("Expected prev link for page 2")
	}
	if response.Meta.Pagination.Links.Next == nil {
		t.Error("Expected next link for page 2")
	}
}
