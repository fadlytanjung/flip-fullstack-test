package handler

import (
	"net/http"

	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	uploadUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/use_case"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// Handler defines the upload handlers
type Handler struct {
	UseCase        uploadUseCase.IUseCase
	CSVValidator   *validator.CSVValidator
	FieldValidator *validator.FieldValidator
}

// NewHandler creates a new upload handler instance
func NewHandler(useCase uploadUseCase.IUseCase) *Handler {
	return &Handler{
		UseCase:        useCase,
		CSVValidator:   validator.NewCSVValidator(),
		FieldValidator: validator.NewFieldValidator(),
	}
}

// Upload handles CSV file uploads
func (h *Handler) Upload(c *fiber.Ctx) error {
	// Parse multipart form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "No file provided",
			Error:   err.Error(),
		})
	}

	// Validate filename
	if err := h.CSVValidator.ValidateFileName(file.Filename); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid filename",
			Error:   err.Error(),
		})
	}

	// Validate file extension
	if err := h.CSVValidator.ValidateFileExtension(file.Filename); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid file type",
			Error:   err.Error(),
		})
	}

	// Validate file header (size, etc)
	if err := h.CSVValidator.ValidateFileHeader(file); err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid file",
			Error:   err.Error(),
		})
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to open file",
			Error:   err.Error(),
		})
	}
	defer src.Close()

	// Parse and store CSV with field validation
	response, err := h.UseCase.ParseAndStoreWithValidation(c.Context(), src, h.FieldValidator)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(schemas.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to process CSV",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

// Clear deletes all transactions
func (h *Handler) Clear(c *fiber.Ctx) error {
	err := h.UseCase.Clear(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(schemas.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to clear transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(schemas.SuccessResponse{
		Status: http.StatusOK,
		Data: fiber.Map{
			"message": "All transactions deleted",
		},
	})
}
