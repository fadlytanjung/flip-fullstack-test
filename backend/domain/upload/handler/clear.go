package handler

import (
	"net/http"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

// Clear deletes all transactions
func (h *Handler) Clear(c *fiber.Ctx) error {
	l := h.Logger.With(
		logger.String("context", ContextName),
		logger.String("method", "Clear"),
	)

	err := h.UseCase.Clear(c.Context())
	if err != nil {
		l.Error("Failed to clear transactions", logger.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: constants.MsgFailedToClearTransactions,
			Error:   err.Error(),
		})
	}

	l.Info("All transactions deleted")

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data: fiber.Map{
			"message": constants.MsgAllTransactionsDeleted,
		},
	})
}

