package handler

import (
	transactionRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/repository"
	uploadRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/repository"
	uploadUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/use_case"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/deps"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

const ContextName = "Domain.Upload.Handler"

// Handler defines the upload handlers
type Handler struct {
	Logger         *logger.Logger
	UseCase        uploadUseCase.IUseCase
	CSVValidator   *validator.CSVValidator
	FieldValidator *validator.FieldValidator
}

// NewHandler creates a new upload handler instance with all dependencies
func NewHandler(d *deps.App) *Handler {
	// Initialize repositories
	uploadRepository := uploadRepo.NewRepository()
	transactionRepository := transactionRepo.NewRepository(d.DB.GetDB())
	
	// Initialize use case
	useCase := uploadUseCase.NewUseCase(uploadRepository, transactionRepository)
	
	return &Handler{
		Logger:         d.Logger,
		UseCase:        useCase,
		CSVValidator:   validator.NewCSVValidator(),
		FieldValidator: validator.NewFieldValidator(),
	}
}

// RegisterApi registers upload API routes
func RegisterApi(d *deps.App) *Handler {
	handler := NewHandler(d)
	
	api := d.Fiber.Group("/api")
	
	api.Post("/upload", handler.Upload)
	api.Delete("/clear", handler.Clear)
	
	return handler
}
