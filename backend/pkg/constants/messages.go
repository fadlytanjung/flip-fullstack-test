package constants

// HTTP Status Messages
const (
	// Generic Messages
	MsgInternalServerError = "Internal server error"
	MsgBadRequest          = "Bad request"
	MsgNotFound            = "Resource not found"
	MsgUnauthorized        = "Unauthorized"
	MsgForbidden           = "Forbidden"
)

// Upload Messages
const (
	MsgUploadSuccess           = "CSV uploaded and processed successfully"
	MsgUploadFailed            = "Failed to process CSV"
	MsgNoFileProvided          = "No file provided"
	MsgInvalidFilename         = "Invalid filename"
	MsgInvalidFileType         = "Invalid file type"
	MsgInvalidFile             = "Invalid file"
	MsgFailedToOpenFile        = "Failed to open file"
	MsgFailedToReadFile        = "Failed to read file"
	MsgNoValidTransactions     = "No valid transactions found in CSV"
	MsgAllTransactionsDeleted  = "All transactions deleted"
	MsgFailedToClearTransactions = "Failed to clear transactions"
)

// Transaction Messages
const (
	MsgBalanceRetrieved         = "Balance retrieved successfully"
	MsgFailedToCalculateBalance = "Failed to calculate balance"
	MsgIssuesRetrieved          = "Issues retrieved successfully"
	MsgFailedToRetrieveIssues   = "Failed to retrieve issues"
	MsgTransactionsRetrieved    = "Transactions retrieved successfully"
	MsgFailedToRetrieveTransactions = "Failed to retrieve transactions"
)

// Validation Messages
const (
	MsgInvalidPagination     = "Invalid pagination parameters"
	MsgInvalidSearchQuery    = "Invalid search query"
	MsgInvalidAmountFilter   = "Invalid amount filter"
	MsgInvalidDateRange      = "Invalid date range"
	MsgInvalidSortField      = "Invalid sort field"
	MsgInvalidSortOrder      = "Invalid sort order"
	MsgInvalidFieldCount     = "Invalid field count"
	MsgInvalidTimestamp      = "Invalid timestamp"
	MsgInvalidName           = "Invalid name"
	MsgInvalidTransactionType = "Invalid transaction type"
	MsgInvalidAmount         = "Invalid amount"
	MsgInvalidStatus         = "Invalid status"
	MsgInvalidDescription    = "Invalid description"
)

// CSV Parsing Messages
const (
	MsgCSVReadError        = "Error reading CSV at line %d: %w"
	MsgCSVInvalidFormat    = "Invalid CSV format at line %d: expected 6 fields, got %d"
	MsgCSVInvalidTimestamp = "Invalid timestamp at line %d: %w"
	MsgCSVInvalidAmount    = "Invalid amount at line %d: %w"
	MsgCSVInvalidType      = "Invalid transaction type at line %d: %s (expected CREDIT or DEBIT)"
	MsgCSVInvalidStatus    = "Invalid status at line %d: %s (expected SUCCESS, FAILED, or PENDING)"
	MsgCSVValidationError  = "Validation error at line %d: %w"
	MsgCSVValidationErrorField = "Validation error at line %d (%s): %w"
)

// Field Validator Error Messages
const (
	ErrMsgFieldCountMismatch     = "expected 6 fields, got %d"
	ErrMsgTimestampFormat        = "invalid timestamp format: must be a Unix epoch integer"
	ErrMsgTimestampRange         = "timestamp out of valid range (must be between 0 and 9999999999)"
	ErrMsgNameEmpty              = "name cannot be empty"
	ErrMsgNameTooLong            = "name exceeds maximum length of 200 characters"
	ErrMsgTypeInvalid            = "type must be either CREDIT or DEBIT"
	ErrMsgAmountFormat           = "invalid amount format: must be a valid number"
	ErrMsgAmountNegative         = "amount cannot be negative"
	ErrMsgStatusInvalid          = "status must be one of: SUCCESS, FAILED, PENDING"
	ErrMsgDescriptionTooLong     = "description exceeds maximum length of 500 characters"
	ErrMsgSearchQueryEmpty       = "search query cannot be empty"
	ErrMsgSearchQueryTooLong     = "search query exceeds maximum length of 200 characters"
	ErrMsgAmountFilterFormat     = "amount filter must be a valid integer"
	ErrMsgDateFormatInvalid      = "date must be in YYYY-MM-DD format"
	ErrMsgDateRangeInvalid       = "start_date cannot be after end_date"
	ErrMsgSortFieldInvalid       = "sort field must be one of: timestamp, name, type, amount, status"
	ErrMsgSortOrderInvalid       = "sort order must be either ASC or DESC"
	ErrMsgPageInvalid            = "page must be greater than 0"
	ErrMsgPageSizeInvalid        = "page_size must be greater than 0"
	ErrMsgPageSizeTooLarge       = "page_size exceeds maximum of 100"
)

// CSV Validator Error Messages
const (
	ErrMsgFilenameEmpty      = "filename cannot be empty"
	ErrMsgFileExtensionInvalid = "file must have .csv extension"
	ErrMsgFileSizeExceeded   = "file size exceeds maximum of %d bytes"
	ErrMsgFileSizeZero       = "file size cannot be zero"
)

