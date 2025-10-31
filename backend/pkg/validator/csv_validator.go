package validator

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// CSVValidator validates CSV files
type CSVValidator struct{}

// NewCSVValidator creates a new CSV validator instance
func NewCSVValidator() *CSVValidator {
	return &CSVValidator{}
}

// ValidateFileExtension checks if file has .csv extension
func (v *CSVValidator) ValidateFileExtension(filename string) error {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != ".csv" {
		return fmt.Errorf("invalid file extension: %s (expected .csv)", ext)
	}
	return nil
}

// ValidateFileHeader checks if file has required CSV headers
func (v *CSVValidator) ValidateFileHeader(header *multipart.FileHeader) error {
	if header == nil {
		return fmt.Errorf("file header is nil")
	}

	if header.Size == 0 {
		return fmt.Errorf("file is empty")
	}

	// Max file size: 10MB
	maxFileSize := int64(10 * 1024 * 1024)
	if header.Size > maxFileSize {
		return fmt.Errorf("file size exceeds maximum allowed size of 10MB")
	}

	return nil
}

// ValidateFileName checks if filename is valid
func (v *CSVValidator) ValidateFileName(filename string) error {
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}

	if len(filename) > 255 {
		return fmt.Errorf("filename is too long (max 255 characters)")
	}

	// Check for invalid characters
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		if strings.Contains(filename, char) {
			return fmt.Errorf("filename contains invalid character: %s", char)
		}
	}

	return nil
}
