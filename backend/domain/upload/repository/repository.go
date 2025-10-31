package repository

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
	"github.com/google/uuid"
)

// IRepository defines the contract for upload repository operations
type IRepository interface {
	// Commands
	ParseCSV(ctx context.Context, file io.Reader) ([]schemas.Transaction, error)
	ParseCSVWithValidation(ctx context.Context, file io.Reader, fieldValidator *validator.FieldValidator) ([]schemas.Transaction, error)

	// Queries
}

// Repository implements IRepository
type Repository struct {
	// Can be extended with dependencies if needed
}

// NewRepository creates a new upload repository instance
func NewRepository() IRepository {
	return &Repository{}
}

// ParseCSV parses a CSV file and returns a slice of Transaction objects (without field validation)
func (r *Repository) ParseCSV(ctx context.Context, file io.Reader) ([]schemas.Transaction, error) {
	// Read all content from the file
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Create CSV reader
	reader := csv.NewReader(bytes.NewReader(content))
	reader.TrimLeadingSpace = true

	var transactions []schemas.Transaction
	seenTransactions := make(map[string]bool) // Track duplicates
	lineNum := 0
	headerSkipped := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV at line %d: %w", lineNum, err)
		}

		lineNum++

		// Skip header row (first row)
		if !headerSkipped {
			headerSkipped = true
			continue
		}

		// Skip empty lines or lines starting with #
		if len(record) == 0 || (len(record) > 0 && strings.HasPrefix(strings.TrimSpace(record[0]), "#")) {
			continue
		}

		// Validate and parse record
		if len(record) < 6 {
			return nil, fmt.Errorf("invalid CSV format at line %d: expected 6 fields, got %d", lineNum, len(record))
		}

		// Parse fields
		timestamp, err := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp at line %d: %w", lineNum, err)
		}

		name := strings.TrimSpace(record[1])
		txType := strings.TrimSpace(record[2])
		txType = strings.ToUpper(txType)

		// Parse amount as float to handle decimal values, then convert to int64 (cents)
		amountFloat, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid amount at line %d: %w", lineNum, err)
		}
		// Convert to cents (multiply by 100 to preserve decimals as integers)
		amount := int64(amountFloat * 100)

		status := strings.TrimSpace(record[4])
		status = strings.ToUpper(status)

		description := strings.TrimSpace(record[5])

		// Validate transaction type
		if txType != string(schemas.TypeCredit) && txType != string(schemas.TypeDebit) {
			return nil, fmt.Errorf("invalid transaction type at line %d: %s (expected CREDIT or DEBIT)", lineNum, txType)
		}

		// Validate status
		validStatus := false
		for _, s := range []string{string(schemas.StatusSuccess), string(schemas.StatusFailed), string(schemas.StatusPending)} {
			if status == s {
				validStatus = true
				break
			}
		}
		if !validStatus {
			return nil, fmt.Errorf("invalid status at line %d: %s (expected SUCCESS, FAILED, or PENDING)", lineNum, status)
		}

		// Create unique key for duplicate detection
		duplicateKey := fmt.Sprintf("%d-%s-%s-%d-%s", timestamp, name, txType, amount, status)

		// Check for duplicates
		if seenTransactions[duplicateKey] {
			continue // Skip duplicate transaction
		}
		seenTransactions[duplicateKey] = true

		// Create transaction object
		transaction := schemas.Transaction{
			ID:          uuid.New().String(),
			Timestamp:   timestamp,
			Name:        name,
			Type:        schemas.TransactionType(txType),
			Amount:      amount,
			Status:      schemas.TransactionStatus(status),
			Description: description,
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("no valid transactions found in CSV")
	}

	return transactions, nil
}

// ParseCSVWithValidation parses a CSV file with field validation
func (r *Repository) ParseCSVWithValidation(ctx context.Context, file io.Reader, fieldValidator *validator.FieldValidator) ([]schemas.Transaction, error) {
	// Read all content from the file
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Create CSV reader
	reader := csv.NewReader(bytes.NewReader(content))
	reader.TrimLeadingSpace = true

	var transactions []schemas.Transaction
	seenTransactions := make(map[string]bool) // Track duplicates
	lineNum := 0
	headerSkipped := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV at line %d: %w", lineNum, err)
		}

		lineNum++

		// Skip header row (first row)
		if !headerSkipped {
			headerSkipped = true
			continue
		}

		// Skip empty lines or lines starting with #
		if len(record) == 0 || (len(record) > 0 && strings.HasPrefix(strings.TrimSpace(record[0]), "#")) {
			continue
		}

		// Validate field count
		if err := fieldValidator.ValidateFieldCount(record); err != nil {
			return nil, fmt.Errorf("validation error at line %d: %w", lineNum, err)
		}

		// Validate each field using field validator
		if err := fieldValidator.ValidateTimestamp(record[0]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (timestamp): %w", lineNum, err)
		}

		if err := fieldValidator.ValidateName(record[1]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (name): %w", lineNum, err)
		}

		if err := fieldValidator.ValidateTransactionType(record[2]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (type): %w", lineNum, err)
		}

		if err := fieldValidator.ValidateAmount(record[3]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (amount): %w", lineNum, err)
		}

		if err := fieldValidator.ValidateStatus(record[4]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (status): %w", lineNum, err)
		}

		if err := fieldValidator.ValidateDescription(record[5]); err != nil {
			return nil, fmt.Errorf("validation error at line %d (description): %w", lineNum, err)
		}

		// Parse fields
		timestamp, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
		name := strings.TrimSpace(record[1])
		txType := strings.ToUpper(strings.TrimSpace(record[2]))
		// Parse amount as float to handle decimal values, then convert to int64 (cents)
		amountFloat, _ := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		amount := int64(amountFloat * 100) // Convert to cents
		status := strings.ToUpper(strings.TrimSpace(record[4]))
		description := strings.TrimSpace(record[5])

		// Create unique key for duplicate detection
		duplicateKey := fmt.Sprintf("%d-%s-%s-%d-%s", timestamp, name, txType, amount, status)

		// Check for duplicates
		if seenTransactions[duplicateKey] {
			continue // Skip duplicate transaction
		}
		seenTransactions[duplicateKey] = true

		// Create transaction object
		transaction := schemas.Transaction{
			ID:          uuid.New().String(),
			Timestamp:   timestamp,
			Name:        name,
			Type:        schemas.TransactionType(txType),
			Amount:      amount,
			Status:      schemas.TransactionStatus(status),
			Description: description,
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("no valid transactions found in CSV")
	}

	return transactions, nil
}