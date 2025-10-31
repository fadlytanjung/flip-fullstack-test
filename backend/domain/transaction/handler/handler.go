package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	transactionUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/use_case"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// Handler defines the transaction handlers
type Handler struct {
	UseCase        transactionUseCase.IUseCase
	FieldValidator *validator.FieldValidator
}

// NewHandler creates a new transaction handler instance
func NewHandler(useCase transactionUseCase.IUseCase) *Handler {
	return &Handler{
		UseCase:        useCase,
		FieldValidator: validator.NewFieldValidator(),
	}
}

// GetBalance returns the calculated balance from successful transactions
func (h *Handler) GetBalance(c *fiber.Ctx) error {
	response, err := h.UseCase.GetBalance(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to calculate balance",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

// GetIssues returns non-successful transactions with pagination, filtering, and sorting
func (h *Handler) GetIssues(c *fiber.Ctx) error {
	// Parse pagination parameters
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if ps := c.Query("page_size"); ps != "" {
		if parsed, err := strconv.Atoi(ps); err == nil && parsed > 0 {
			pageSize = parsed
		}
	}

	// Limit page size to 100
	if pageSize > 100 {
		pageSize = 100
	}

	// Validate pagination
	if err := h.FieldValidator.ValidatePaginationParams(page, pageSize); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid pagination parameters",
			Error:   err.Error(),
		})
	}

	// Parse filter parameters
	filters := schemas.TransactionFilters{
		Status:      strings.ToUpper(c.Query("status")),
		Type:        strings.ToUpper(c.Query("type")),
		SearchQuery: c.Query("search"),
		StartDate:   c.Query("start_date"),
		EndDate:     c.Query("end_date"),
	}

	// Validate search query
	if filters.SearchQuery != "" {
		if err := h.FieldValidator.ValidateSearchQuery(filters.SearchQuery); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid search query",
				Error:   err.Error(),
			})
		}
	}

	// Parse amount filter if provided
	if amountStr := c.Query("amount"); amountStr != "" {
		if err := h.FieldValidator.ValidateAmountFilter(amountStr); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid amount filter",
				Error:   err.Error(),
			})
		}
		if amount, err := strconv.ParseInt(amountStr, 10, 64); err == nil {
			filters.Amount = amount
		}
	}

	// Validate date range
	if err := h.FieldValidator.ValidateDateRange(filters.StartDate, filters.EndDate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid date range",
			Error:   err.Error(),
		})
	}

	// Parse sort parameters (optional - no defaults)
	sortBy := c.Query("sort_by", "")
	sortOrder := c.Query("sort_order", "")

	// Only validate sort parameters if they are provided
	if sortBy != "" {
		if err := h.FieldValidator.ValidateSortField(sortBy); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid sort field",
				Error:   err.Error(),
			})
		}
	}

	if sortOrder != "" {
		if err := h.FieldValidator.ValidateSortOrder(sortOrder); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid sort order",
				Error:   err.Error(),
			})
		}
	}

	sort := schemas.TransactionSort{
		By:    sortBy,
		Order: strings.ToUpper(sortOrder),
	}

	response, err := h.UseCase.GetIssuesWithFiltersAndSort(c.Context(), page, pageSize, filters, sort)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve issues",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

// GetTransactions returns all transactions with pagination, filtering, and sorting
func (h *Handler) GetTransactions(c *fiber.Ctx) error {
	// Parse pagination parameters
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if ps := c.Query("page_size"); ps != "" {
		if parsed, err := strconv.Atoi(ps); err == nil && parsed > 0 {
			pageSize = parsed
		}
	}

	// Limit page size to 100
	if pageSize > 100 {
		pageSize = 100
	}

	// Validate pagination
	if err := h.FieldValidator.ValidatePaginationParams(page, pageSize); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid pagination parameters",
			Error:   err.Error(),
		})
	}

	// Parse filter parameters
	filters := schemas.TransactionFilters{
		Status:      strings.ToUpper(c.Query("status")),
		Type:        strings.ToUpper(c.Query("type")),
		SearchQuery: c.Query("search"),
		StartDate:   c.Query("start_date"),
		EndDate:     c.Query("end_date"),
	}

	// Validate search query
	if filters.SearchQuery != "" {
		if err := h.FieldValidator.ValidateSearchQuery(filters.SearchQuery); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid search query",
				Error:   err.Error(),
			})
		}
	}

	// Parse amount filter if provided
	if amountStr := c.Query("amount"); amountStr != "" {
		if err := h.FieldValidator.ValidateAmountFilter(amountStr); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid amount filter",
				Error:   err.Error(),
			})
		}
		if amount, err := strconv.ParseInt(amountStr, 10, 64); err == nil {
			filters.Amount = amount
		}
	}

	// Validate date range
	if err := h.FieldValidator.ValidateDateRange(filters.StartDate, filters.EndDate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid date range",
			Error:   err.Error(),
		})
	}

	// Parse sort parameters (optional - no defaults)
	sortBy := c.Query("sort_by", "")
	sortOrder := c.Query("sort_order", "")

	// Only validate sort parameters if they are provided
	if sortBy != "" {
		if err := h.FieldValidator.ValidateSortField(sortBy); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid sort field",
				Error:   err.Error(),
			})
		}
	}

	if sortOrder != "" {
		if err := h.FieldValidator.ValidateSortOrder(sortOrder); err != nil {
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid sort order",
				Error:   err.Error(),
			})
		}
	}

	sort := schemas.TransactionSort{
		By:    sortBy,
		Order: strings.ToUpper(sortOrder),
	}

	response, err := h.UseCase.GetAllWithFiltersAndSort(c.Context(), page, pageSize, filters, sort)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}
