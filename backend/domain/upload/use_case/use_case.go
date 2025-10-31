package use_case

import (
	"context"
	"io"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/repository"
	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	uploadRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/repository"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
)

// IUseCase defines the contract for upload use case operations
type IUseCase interface {
	ParseAndStore(ctx context.Context, file io.Reader) (*schemas.UploadResponse, error)
	ParseAndStoreWithValidation(ctx context.Context, file io.Reader, fieldValidator *validator.FieldValidator) (*schemas.UploadResponse, error)
	Clear(ctx context.Context) error
}

// UseCase implements IUseCase
type UseCase struct {
	uploadRepo      uploadRepo.IRepository
	transactionRepo repository.IRepository
}

// NewUseCase creates a new upload use case instance
func NewUseCase(uploadRepo uploadRepo.IRepository, transactionRepo repository.IRepository) IUseCase {
	return &UseCase{
		uploadRepo:      uploadRepo,
		transactionRepo: transactionRepo,
	}
}

// ParseAndStore parses CSV file and stores transactions (without field validation)
func (uc *UseCase) ParseAndStore(ctx context.Context, file io.Reader) (*schemas.UploadResponse, error) {
	// Parse CSV
	transactions, err := uc.uploadRepo.ParseCSV(ctx, file)
	if err != nil {
		return nil, err
	}

	// Store transactions in database
	err = uc.transactionRepo.CreateBatch(ctx, transactions)
	if err != nil {
		return nil, err
	}

	// Count transactions by status
	successCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusSuccess)
	failedCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusFailed)
	pendingCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusPending)

	return &schemas.UploadResponse{
		Message:        "CSV uploaded and processed successfully",
		TotalRecords:   len(transactions),
		SuccessRecords: int(successCount),
		FailedRecords:  int(failedCount),
		PendingRecords: int(pendingCount),
	}, nil
}

// ParseAndStoreWithValidation parses CSV file with field validation and stores transactions
func (uc *UseCase) ParseAndStoreWithValidation(ctx context.Context, file io.Reader, fieldValidator *validator.FieldValidator) (*schemas.UploadResponse, error) {
	// Parse CSV with field validation
	transactions, err := uc.uploadRepo.ParseCSVWithValidation(ctx, file, fieldValidator)
	if err != nil {
		return nil, err
	}

	// Store transactions in database
	err = uc.transactionRepo.CreateBatch(ctx, transactions)
	if err != nil {
		return nil, err
	}

	// Count transactions by status
	successCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusSuccess)
	failedCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusFailed)
	pendingCount, _ := uc.transactionRepo.CountByStatus(ctx, schemas.StatusPending)

	return &schemas.UploadResponse{
		Message:        "CSV uploaded and processed successfully",
		TotalRecords:   len(transactions),
		SuccessRecords: int(successCount),
		FailedRecords:  int(failedCount),
		PendingRecords: int(pendingCount),
	}, nil
}

// Clear deletes all transactions from database
func (uc *UseCase) Clear(ctx context.Context) error {
	return uc.transactionRepo.DeleteAll(ctx)
}
