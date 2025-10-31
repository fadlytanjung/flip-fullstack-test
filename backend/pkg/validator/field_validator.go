package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FieldValidator validates individual transaction fields
type FieldValidator struct{}

// NewFieldValidator creates a new field validator instance
func NewFieldValidator() *FieldValidator {
	return &FieldValidator{}
}

// ValidateTimestamp validates if timestamp is a valid Unix epoch
func (v *FieldValidator) ValidateTimestamp(timestamp string) error {
	timestamp = strings.TrimSpace(timestamp)
	if timestamp == "" {
		return fmt.Errorf("timestamp is required")
	}

	_, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: must be a Unix epoch integer")
	}

	return nil
}

// ValidateName validates if name is not empty
func (v *FieldValidator) ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("name is required")
	}

	if len(name) > 255 {
		return fmt.Errorf("name is too long (max 255 characters)")
	}

	return nil
}

// ValidateTransactionType validates if type is CREDIT or DEBIT
func (v *FieldValidator) ValidateTransactionType(txType string) error {
	txType = strings.TrimSpace(strings.ToUpper(txType))
	if txType == "" {
		return fmt.Errorf("transaction type is required")
	}

	if txType != "CREDIT" && txType != "DEBIT" {
		return fmt.Errorf("invalid transaction type: %s (expected CREDIT or DEBIT)", txType)
	}

	return nil
}

// ValidateAmount validates if amount is a valid integer or decimal number
func (v *FieldValidator) ValidateAmount(amount string) error {
	amount = strings.TrimSpace(amount)
	if amount == "" {
		return fmt.Errorf("amount is required")
	}

	// Try to parse as float first (to handle decimal values like 123.45)
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount format: must be a valid number")
	}

	if parsedAmount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	return nil
}

// ValidateStatus validates if status is SUCCESS, FAILED, or PENDING
func (v *FieldValidator) ValidateStatus(status string) error {
	status = strings.TrimSpace(strings.ToUpper(status))
	if status == "" {
		return fmt.Errorf("status is required")
	}

	validStatuses := []string{"SUCCESS", "FAILED", "PENDING"}
	for _, s := range validStatuses {
		if status == s {
			return nil
		}
	}

	return fmt.Errorf("invalid status: %s (expected SUCCESS, FAILED, or PENDING)", status)
}

// ValidateDescription validates if description is not too long
func (v *FieldValidator) ValidateDescription(description string) error {
	description = strings.TrimSpace(description)

	if len(description) > 500 {
		return fmt.Errorf("description is too long (max 500 characters)")
	}

	return nil
}

// ValidateFieldCount validates if CSV record has exactly 6 fields
func (v *FieldValidator) ValidateFieldCount(fields []string) error {
	if len(fields) < 6 {
		return fmt.Errorf("invalid CSV format: expected 6 fields, got %d", len(fields))
	}
	return nil
}

// ValidateSearchQuery validates search query format
func (v *FieldValidator) ValidateSearchQuery(query string) error {
	if len(query) > 255 {
		return fmt.Errorf("search query is too long (max 255 characters)")
	}

	// Check for SQL injection patterns
	sqlInjectionPatterns := []string{
		"'; DROP",
		"'; DELETE",
		"'; UPDATE",
		"'; INSERT",
		"--",
		"/*",
		"*/",
	}

	upperQuery := strings.ToUpper(query)
	for _, pattern := range sqlInjectionPatterns {
		if strings.Contains(upperQuery, pattern) {
			return fmt.Errorf("invalid search query: contains suspicious pattern")
		}
	}

	return nil
}

// ValidateSortField validates if sort field is allowed
func (v *FieldValidator) ValidateSortField(field string) error {
	allowedFields := []string{
		"timestamp",
		"amount",
		"name",
		"status",
		"type",
		"description",
		"created_at",
	}

	field = strings.ToLower(strings.TrimSpace(field))
	for _, allowed := range allowedFields {
		if field == allowed {
			return nil
		}
	}

	return fmt.Errorf("invalid sort field: %s", field)
}

// ValidateSortOrder validates if sort order is ASC or DESC
func (v *FieldValidator) ValidateSortOrder(order string) error {
	order = strings.ToUpper(strings.TrimSpace(order))
	if order == "" {
		return nil // Default to DESC if not provided
	}

	if order != "ASC" && order != "DESC" {
		return fmt.Errorf("invalid sort order: %s (expected ASC or DESC)", order)
	}

	return nil
}

// ValidatePaginationParams validates pagination parameters
func (v *FieldValidator) ValidatePaginationParams(page, pageSize int) error {
	if page < 1 {
		return fmt.Errorf("page must be >= 1")
	}

	if pageSize < 1 || pageSize > 100 {
		return fmt.Errorf("page_size must be between 1 and 100")
	}

	return nil
}

// ValidateAmount validates if filter amount is valid
func (v *FieldValidator) ValidateAmountFilter(amount string) error {
	if amount == "" {
		return nil // Optional filter
	}

	amount = strings.TrimSpace(amount)
	_, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid amount filter: must be an integer")
	}

	return nil
}

// ValidateDateRange validates if date range is valid
func (v *FieldValidator) ValidateDateRange(startDate, endDate string) error {
	// Optional filters
	if startDate == "" && endDate == "" {
		return nil
	}

	// Simple ISO date validation
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	if startDate != "" && !dateRegex.MatchString(startDate) {
		return fmt.Errorf("invalid start_date format: expected YYYY-MM-DD")
	}

	if endDate != "" && !dateRegex.MatchString(endDate) {
		return fmt.Errorf("invalid end_date format: expected YYYY-MM-DD")
	}

	return nil
}
