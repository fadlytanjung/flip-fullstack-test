package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

// GetTransactions returns all transactions with pagination, filtering, and sorting
func (h *Handler) GetTransactions(c *fiber.Ctx) error {
	l := h.Logger.With(
		logger.String("context", ContextName),
		logger.String("method", "GetTransactions"),
	)

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
		l.Warn("Invalid pagination parameters", logger.Error(err))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgInvalidPagination,
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
			l.Warn("Invalid search query", logger.Error(err))
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: constants.MsgInvalidSearchQuery,
				Error:   err.Error(),
			})
		}
	}

	// Parse amount filter if provided
	if amountStr := c.Query("amount"); amountStr != "" {
		if err := h.FieldValidator.ValidateAmountFilter(amountStr); err != nil {
			l.Warn("Invalid amount filter", logger.Error(err))
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: constants.MsgInvalidAmountFilter,
				Error:   err.Error(),
			})
		}
		if amount, err := strconv.ParseInt(amountStr, 10, 64); err == nil {
			filters.Amount = amount
		}
	}

	// Validate date range
	if err := h.FieldValidator.ValidateDateRange(filters.StartDate, filters.EndDate); err != nil {
		l.Warn("Invalid date range", logger.Error(err))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgInvalidDateRange,
			Error:   err.Error(),
		})
	}

	// Parse sort parameters (optional - no defaults)
	sortBy := c.Query("sort_by", "")
	sortOrder := c.Query("sort_order", "")

	// Only validate sort parameters if they are provided
	if sortBy != "" {
		if err := h.FieldValidator.ValidateSortField(sortBy); err != nil {
			l.Warn("Invalid sort field", logger.Error(err))
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: constants.MsgInvalidSortField,
				Error:   err.Error(),
			})
		}
	}

	if sortOrder != "" {
		if err := h.FieldValidator.ValidateSortOrder(sortOrder); err != nil {
			l.Warn("Invalid sort order", logger.Error(err))
			return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: constants.MsgInvalidSortOrder,
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
		l.Error("Failed to retrieve transactions", logger.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: constants.MsgFailedToRetrieveTransactions,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

