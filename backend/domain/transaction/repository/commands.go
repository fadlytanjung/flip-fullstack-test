package repository

import (
	"context"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
)

// Create creates a single transaction record
func (r *Repository) Create(ctx context.Context, transaction *schemas.Transaction) error {
	return r.DB.WithContext(ctx).Create(transaction).Error
}

// CreateBatch creates multiple transaction records
func (r *Repository) CreateBatch(ctx context.Context, transactions []schemas.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).CreateInBatches(transactions, 100).Error
}

// DeleteAll deletes all transaction records
func (r *Repository) DeleteAll(ctx context.Context) error {
	return r.DB.WithContext(ctx).Exec("DELETE FROM transactions").Error
}
