package repository

import (
	"bytes"
	"context"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
)

// TestParseCSVValid tests parsing a valid CSV
func TestParseCSVValid(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
1624512883,COMPANY A,CREDIT,12000000,SUCCESS,salary`

	repo := NewRepository()
	ctx := context.Background()
	transactions, err := repo.ParseCSV(ctx, bytes.NewBufferString(csvContent))
	if err != nil {
		t.Fatalf("ParseCSV failed: %v", err)
	}

	if len(transactions) != 3 {
		t.Errorf("Expected 3 transactions, got %d", len(transactions))
	}

	// Verify first transaction
	if transactions[0].Name != "JOHN DOE" {
		t.Errorf("Expected name 'JOHN DOE', got %s", transactions[0].Name)
	}

	if transactions[0].Type != schemas.TypeDebit {
		t.Errorf("Expected type DEBIT, got %s", transactions[0].Type)
	}

	if transactions[0].Amount != 250000 {
		t.Errorf("Expected amount 250000, got %d", transactions[0].Amount)
	}

	if transactions[0].Status != schemas.StatusSuccess {
		t.Errorf("Expected status SUCCESS, got %s", transactions[0].Status)
	}
}

// TestParseCSVWithSpaces tests CSV parsing with extra spaces
func TestParseCSVWithSpaces(t *testing.T) {
	csvContent := `1624507883 , JOHN DOE, DEBIT, 250000 , SUCCESS, restaurant
1624608050 , E-COMMERCE A, DEBIT, 150000 , FAILED, clothes`

	repo := NewRepository()
	ctx := context.Background()
	transactions, err := repo.ParseCSV(ctx, bytes.NewBufferString(csvContent))
	if err != nil {
		t.Fatalf("ParseCSV failed: %v", err)
	}

	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}

	// Verify spaces are trimmed
	if transactions[0].Name != "JOHN DOE" {
		t.Errorf("Expected name 'JOHN DOE', got '%s'", transactions[0].Name)
	}

	if transactions[0].Type != schemas.TypeDebit {
		t.Errorf("Expected type DEBIT, got %s", transactions[0].Type)
	}
}

// TestParseCSVWithValidation tests CSV parsing with field validation
func TestParseCSVWithValidation(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
1624512883,COMPANY A,CREDIT,12000000,SUCCESS,salary`

	fieldValidator := validator.NewFieldValidator()
	repo := NewRepository()
	ctx := context.Background()

	transactions, err := repo.ParseCSVWithValidation(ctx, bytes.NewBufferString(csvContent), fieldValidator)

	if err != nil {
		t.Fatalf("ParseCSVWithValidation failed: %v", err)
	}

	if len(transactions) != 3 {
		t.Errorf("Expected 3 valid transactions, got %d", len(transactions))
	}
}

// TestParseCSVWithInvalidData tests CSV parsing with invalid data
func TestParseCSVWithInvalidData(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,INVALID_TYPE,250000,SUCCESS,restaurant
invalid_timestamp,E-COMMERCE A,DEBIT,150000,FAILED,clothes`

	fieldValidator := validator.NewFieldValidator()
	repo := NewRepository()
	ctx := context.Background()

	_, err := repo.ParseCSVWithValidation(ctx, bytes.NewBufferString(csvContent), fieldValidator)

	// Should error on invalid type
	if err == nil {
		t.Error("Expected error for invalid data, got nil")
	}
}

// TestParseCSVEmptyFile tests parsing an empty CSV
func TestParseCSVEmptyFile(t *testing.T) {
	csvContent := ""

	repo := NewRepository()
	ctx := context.Background()
	_, err := repo.ParseCSV(ctx, bytes.NewBufferString(csvContent))

	// Empty file should return error
	if err == nil {
		t.Error("Expected error for empty CSV")
	}
}

// TestParseCSVMissingFields tests CSV with missing fields
func TestParseCSVMissingFields(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,250000`

	fieldValidator := validator.NewFieldValidator()
	repo := NewRepository()
	ctx := context.Background()

	_, err := repo.ParseCSVWithValidation(ctx, bytes.NewBufferString(csvContent), fieldValidator)

	if err == nil {
		t.Error("Expected error for missing fields")
	}

	if !strings.Contains(err.Error(), "6 fields") {
		t.Errorf("Expected field count error, got: %v", err)
	}
}

// TestParseCSVWithAmount tests amount field validation
func TestParseCSVWithAmount(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,-100,SUCCESS,restaurant`

	fieldValidator := validator.NewFieldValidator()
	repo := NewRepository()
	ctx := context.Background()

	_, err := repo.ParseCSVWithValidation(ctx, bytes.NewBufferString(csvContent), fieldValidator)

	// Should reject negative amount
	if err == nil {
		t.Error("Expected error for negative amount")
	}

	if !strings.Contains(err.Error(), "negative") {
		t.Errorf("Expected error about negative amount, got: %v", err)
	}
}

// Helper function to create a test multipart file
func createTestMultipartFile(t *testing.T, content string) *multipart.FileHeader {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, err = part.Write([]byte(content))
	if err != nil {
		t.Fatalf("Failed to write to form file: %v", err)
	}

	writer.Close()

	// Parse the multipart data
	reader := multipart.NewReader(body, writer.Boundary())
	form, err := reader.ReadForm(10 << 20) // 10MB max
	if err != nil {
		t.Fatalf("Failed to read form: %v", err)
	}

	if len(form.File["file"]) == 0 {
		t.Fatal("No file in form")
	}

	return form.File["file"][0]
}

// TestParseCSVFromMultipart tests parsing CSV from multipart file
func TestParseCSVFromMultipart(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes`

	header := createTestMultipartFile(t, csvContent)

	file, err := header.Open()
	if err != nil {
		t.Fatalf("Failed to open multipart file: %v", err)
	}
	defer file.Close()

	repo := NewRepository()
	ctx := context.Background()
	transactions, err := repo.ParseCSV(ctx, file)
	if err != nil {
		t.Fatalf("ParseCSV from multipart failed: %v", err)
	}

	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions from multipart file, got %d", len(transactions))
	}
}

// TestParseCSVAllFields tests parsing all required fields correctly
func TestParseCSVAllFields(t *testing.T) {
	csvContent := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant payment
1624608050,E-COMMERCE A,CREDIT,150000,PENDING,online purchase`

	repo := NewRepository()
	ctx := context.Background()
	transactions, err := repo.ParseCSV(ctx, bytes.NewBufferString(csvContent))
	if err != nil {
		t.Fatalf("ParseCSV failed: %v", err)
	}

	// Check first transaction
	tx1 := transactions[0]
	if tx1.Timestamp != 1624507883 {
		t.Errorf("Expected timestamp 1624507883, got %d", tx1.Timestamp)
	}
	if tx1.Name != "JOHN DOE" {
		t.Errorf("Expected name 'JOHN DOE', got %s", tx1.Name)
	}
	if tx1.Type != schemas.TypeDebit {
		t.Errorf("Expected type DEBIT, got %s", tx1.Type)
	}
	if tx1.Amount != 250000 {
		t.Errorf("Expected amount 250000, got %d", tx1.Amount)
	}
	if tx1.Status != schemas.StatusSuccess {
		t.Errorf("Expected status SUCCESS, got %s", tx1.Status)
	}
	if tx1.Description != "restaurant payment" {
		t.Errorf("Expected description 'restaurant payment', got %s", tx1.Description)
	}

	// Check second transaction
	tx2 := transactions[1]
	if tx2.Type != schemas.TypeCredit {
		t.Errorf("Expected type CREDIT, got %s", tx2.Type)
	}
	if tx2.Status != schemas.StatusPending {
		t.Errorf("Expected status PENDING, got %s", tx2.Status)
	}
}
