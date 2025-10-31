package validator

import (
	"testing"
)

// TestValidateFileExtension tests file extension validation
func TestValidateFileExtension(t *testing.T) {
	validator := NewCSVValidator()

	tests := []struct {
		name      string
		filename  string
		shouldErr bool
	}{
		{
			name:      "valid csv extension",
			filename:  "transactions.csv",
			shouldErr: false,
		},
		{
			name:      "valid csv with uppercase",
			filename:  "TRANSACTIONS.CSV",
			shouldErr: false,
		},
		{
			name:      "invalid txt extension",
			filename:  "transactions.txt",
			shouldErr: true,
		},
		{
			name:      "invalid xlsx extension",
			filename:  "transactions.xlsx",
			shouldErr: true,
		},
		{
			name:      "no extension",
			filename:  "transactions",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateFileExtension(tc.filename)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for filename: %s", tc.filename)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for filename: %s, err: %v", tc.filename, err)
			}
		})
	}
}

// TestValidateFileName tests filename validation
func TestValidateFileName(t *testing.T) {
	validator := NewCSVValidator()

	tests := []struct {
		name      string
		filename  string
		shouldErr bool
	}{
		{
			name:      "valid filename",
			filename:  "transactions.csv",
			shouldErr: false,
		},
		{
			name:      "valid filename with numbers",
			filename:  "transactions_2024.csv",
			shouldErr: false,
		},
		{
			name:      "empty filename",
			filename:  "",
			shouldErr: true,
		},
		{
			name:      "filename with forward slash",
			filename:  "folder/transactions.csv",
			shouldErr: true,
		},
		{
			name:      "filename with backslash",
			filename:  "folder\\transactions.csv",
			shouldErr: true,
		},
		{
			name:      "filename with asterisk",
			filename:  "trans*.csv",
			shouldErr: true,
		},
		{
			name:      "filename with question mark",
			filename:  "trans?.csv",
			shouldErr: true,
		},
		{
			name:      "very long filename",
			filename:  string(make([]byte, 256)) + "a",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateFileName(tc.filename)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for filename: %s", tc.filename)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for filename: %s, err: %v", tc.filename, err)
			}
		})
	}
}

// TestValidateTimestamp tests timestamp validation
func TestValidateTimestamp(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		timestamp string
		shouldErr bool
	}{
		{
			name:      "valid timestamp",
			timestamp: "1624507883",
			shouldErr: false,
		},
		{
			name:      "timestamp with spaces",
			timestamp: "  1624507883  ",
			shouldErr: false,
		},
		{
			name:      "empty timestamp",
			timestamp: "",
			shouldErr: true,
		},
		{
			name:      "non-numeric timestamp",
			timestamp: "abc",
			shouldErr: true,
		},
		{
			name:      "timestamp with decimals",
			timestamp: "1624507883.5",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateTimestamp(tc.timestamp)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for timestamp: %s", tc.timestamp)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for timestamp: %s, err: %v", tc.timestamp, err)
			}
		})
	}
}

// TestValidateName tests name validation
func TestValidateName(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		nameValue string
		shouldErr bool
	}{
		{
			name:      "valid name",
			nameValue: "JOHN DOE",
			shouldErr: false,
		},
		{
			name:      "valid name with numbers",
			nameValue: "Company 123",
			shouldErr: false,
		},
		{
			name:      "empty name",
			nameValue: "",
			shouldErr: true,
		},
		{
			name:      "name with spaces only",
			nameValue: "   ",
			shouldErr: true,
		},
		{
			name:      "very long name",
			nameValue: string(make([]byte, 256)),
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateName(tc.nameValue)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for name: %s", tc.nameValue)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for name: %s, err: %v", tc.nameValue, err)
			}
		})
	}
}

// TestValidateTransactionType tests transaction type validation
func TestValidateTransactionType(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		txType    string
		shouldErr bool
	}{
		{
			name:      "valid credit type",
			txType:    "CREDIT",
			shouldErr: false,
		},
		{
			name:      "valid debit type",
			txType:    "DEBIT",
			shouldErr: false,
		},
		{
			name:      "lowercase credit",
			txType:    "credit",
			shouldErr: false,
		},
		{
			name:      "empty type",
			txType:    "",
			shouldErr: true,
		},
		{
			name:      "invalid type",
			txType:    "TRANSFER",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateTransactionType(tc.txType)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for type: %s", tc.txType)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for type: %s, err: %v", tc.txType, err)
			}
		})
	}
}

// TestValidateAmount tests amount validation
func TestValidateAmount(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		amount    string
		shouldErr bool
	}{
		{
			name:      "valid amount",
			amount:    "250000",
			shouldErr: false,
		},
		{
			name:      "valid large amount",
			amount:    "12000000",
			shouldErr: false,
		},
		{
			name:      "amount with spaces",
			amount:    "  250000  ",
			shouldErr: false,
		},
		{
			name:      "zero amount",
			amount:    "0",
			shouldErr: false,
		},
		{
			name:      "negative amount",
			amount:    "-250000",
			shouldErr: true,
		},
		{
			name:      "empty amount",
			amount:    "",
			shouldErr: true,
		},
		{
			name:      "non-numeric amount",
			amount:    "abc",
			shouldErr: true,
		},
		{
			name:      "decimal amount",
			amount:    "250000.50",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateAmount(tc.amount)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for amount: %s", tc.amount)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for amount: %s, err: %v", tc.amount, err)
			}
		})
	}
}

// TestValidateStatus tests status validation
func TestValidateStatus(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		status    string
		shouldErr bool
	}{
		{
			name:      "valid success status",
			status:    "SUCCESS",
			shouldErr: false,
		},
		{
			name:      "valid failed status",
			status:    "FAILED",
			shouldErr: false,
		},
		{
			name:      "valid pending status",
			status:    "PENDING",
			shouldErr: false,
		},
		{
			name:      "lowercase success",
			status:    "success",
			shouldErr: false,
		},
		{
			name:      "empty status",
			status:    "",
			shouldErr: true,
		},
		{
			name:      "invalid status",
			status:    "COMPLETED",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateStatus(tc.status)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for status: %s", tc.status)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for status: %s, err: %v", tc.status, err)
			}
		})
	}
}

// TestValidateDescription tests description validation
func TestValidateDescription(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name        string
		description string
		shouldErr   bool
	}{
		{
			name:        "valid description",
			description: "restaurant payment",
			shouldErr:   false,
		},
		{
			name:        "empty description",
			description: "",
			shouldErr:   false, // Description is optional
		},
		{
			name:        "very long description",
			description: string(make([]byte, 501)),
			shouldErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateDescription(tc.description)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for description: %s", tc.description)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for description: %s, err: %v", tc.description, err)
			}
		})
	}
}

// TestValidateSortField tests sort field validation
func TestValidateSortField(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		field     string
		shouldErr bool
	}{
		{
			name:      "valid timestamp field",
			field:     "timestamp",
			shouldErr: false,
		},
		{
			name:      "valid amount field",
			field:     "amount",
			shouldErr: false,
		},
		{
			name:      "valid status field",
			field:     "status",
			shouldErr: false,
		},
		{
			name:      "uppercase field",
			field:     "TIMESTAMP",
			shouldErr: false,
		},
		{
			name:      "invalid field",
			field:     "invalid_field",
			shouldErr: true,
		},
		{
			name:      "empty field",
			field:     "",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateSortField(tc.field)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for field: %s", tc.field)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for field: %s, err: %v", tc.field, err)
			}
		})
	}
}

// TestValidateSortOrder tests sort order validation
func TestValidateSortOrder(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		order     string
		shouldErr bool
	}{
		{
			name:      "valid ASC order",
			order:     "ASC",
			shouldErr: false,
		},
		{
			name:      "valid DESC order",
			order:     "DESC",
			shouldErr: false,
		},
		{
			name:      "lowercase asc",
			order:     "asc",
			shouldErr: false,
		},
		{
			name:      "empty order (default)",
			order:     "",
			shouldErr: false,
		},
		{
			name:      "invalid order",
			order:     "INVALID",
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateSortOrder(tc.order)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for order: %s", tc.order)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for order: %s, err: %v", tc.order, err)
			}
		})
	}
}

// TestValidatePaginationParams tests pagination parameter validation
func TestValidatePaginationParams(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		page      int
		pageSize  int
		shouldErr bool
	}{
		{
			name:      "valid parameters",
			page:      1,
			pageSize:  10,
			shouldErr: false,
		},
		{
			name:      "valid max page size",
			page:      1,
			pageSize:  100,
			shouldErr: false,
		},
		{
			name:      "page 0",
			page:      0,
			pageSize:  10,
			shouldErr: true,
		},
		{
			name:      "negative page",
			page:      -1,
			pageSize:  10,
			shouldErr: true,
		},
		{
			name:      "page size 0",
			page:      1,
			pageSize:  0,
			shouldErr: true,
		},
		{
			name:      "page size too large",
			page:      1,
			pageSize:  101,
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidatePaginationParams(tc.page, tc.pageSize)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for page=%d, pageSize=%d", tc.page, tc.pageSize)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for page=%d, pageSize=%d, err: %v", tc.page, tc.pageSize, err)
			}
		})
	}
}

// TestValidateSearchQuery tests search query validation
func TestValidateSearchQuery(t *testing.T) {
	validator := NewFieldValidator()

	tests := []struct {
		name      string
		query     string
		shouldErr bool
	}{
		{
			name:      "valid search query",
			query:     "restaurant",
			shouldErr: false,
		},
		{
			name:      "valid query with spaces",
			query:     "e-commerce payment",
			shouldErr: false,
		},
		{
			name:      "empty query",
			query:     "",
			shouldErr: false, // Optional
		},
		{
			name:      "query with SQL injection attempt",
			query:     "'; DROP TABLE",
			shouldErr: true,
		},
		{
			name:      "query with comment",
			query:     "test -- comment",
			shouldErr: true,
		},
		{
			name:      "very long query",
			query:     string(make([]byte, 256)),
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateSearchQuery(tc.query)
			if tc.shouldErr && err == nil {
				t.Errorf("Expected error for query: %s", tc.query)
			}
			if !tc.shouldErr && err != nil {
				t.Errorf("Unexpected error for query: %s, err: %v", tc.query, err)
			}
		})
	}
}
