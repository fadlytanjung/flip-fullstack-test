package handler

import (
	"net/http"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

// GetBalance returns the calculated balance from successful transactions
func (h *Handler) GetBalance(c *fiber.Ctx) error {
	l := h.Logger.With(
		logger.String("context", ContextName),
		logger.String("method", "GetBalance"),
	)

	response, err := h.UseCase.GetBalance(c.Context())
	if err != nil {
		l.Error("Failed to calculate balance", logger.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: constants.MsgFailedToCalculateBalance,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

