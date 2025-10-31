package handler

import (
	"net/http"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

// Upload handles CSV file uploads
func (h *Handler) Upload(c *fiber.Ctx) error {
	l := h.Logger.With(
		logger.String("context", ContextName),
		logger.String("method", "Upload"),
	)

	// Parse multipart form
	file, err := c.FormFile("file")
	if err != nil {
		l.Warn("No file provided", logger.Error(err))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgNoFileProvided,
			Error:   err.Error(),
		})
	}

	// Validate filename
	if err := h.CSVValidator.ValidateFileName(file.Filename); err != nil {
		l.Warn("Invalid filename", logger.Error(err), logger.String("filename", file.Filename))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgInvalidFilename,
			Error:   err.Error(),
		})
	}

	// Validate file extension
	if err := h.CSVValidator.ValidateFileExtension(file.Filename); err != nil {
		l.Warn("Invalid file type", logger.Error(err), logger.String("filename", file.Filename))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgInvalidFileType,
			Error:   err.Error(),
		})
	}

	// Validate file header (size, etc)
	if err := h.CSVValidator.ValidateFileHeader(file); err != nil {
		l.Warn("Invalid file", logger.Error(err), logger.Int64("size", file.Size))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgInvalidFile,
			Error:   err.Error(),
		})
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		l.Error("Failed to open file", logger.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: constants.MsgFailedToOpenFile,
			Error:   err.Error(),
		})
	}
	defer src.Close()

	// Parse and store CSV with field validation
	response, err := h.UseCase.ParseAndStoreWithValidation(c.Context(), src, h.FieldValidator)
	if err != nil {
		l.Error("Failed to process CSV", logger.Error(err))
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: constants.MsgUploadFailed,
			Error:   err.Error(),
		})
	}

	l.Info("CSV uploaded successfully",
		logger.Int("total_records", response.TotalRecords),
		logger.Int("success_records", response.SuccessRecords),
		logger.Int("failed_records", response.FailedRecords),
		logger.Int("pending_records", response.PendingRecords),
	)

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

