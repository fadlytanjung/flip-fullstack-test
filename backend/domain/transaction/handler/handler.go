package handler

import (
	transactionRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/repository"
	transactionUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/use_case"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/deps"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/validator"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
)

const ContextName = "Domain.Transaction.Handler"

// Handler defines the transaction handlers
type Handler struct {
	Logger         *logger.Logger
	UseCase        transactionUseCase.IUseCase
	FieldValidator *validator.FieldValidator
}

// NewHandler creates a new transaction handler instance with all dependencies
func NewHandler(d *deps.App) *Handler {
	// Initialize repository
	repository := transactionRepo.NewRepository(d.DB.GetDB())
	
	// Initialize use case
	useCase := transactionUseCase.NewUseCase(repository)
	
	return &Handler{
		Logger:         d.Logger,
		UseCase:        useCase,
		FieldValidator: validator.NewFieldValidator(),
	}
}

// RegisterApi registers transaction API routes
func RegisterApi(d *deps.App) *Handler {
	handler := NewHandler(d)
	
	api := d.Fiber.Group("/api")
	
	api.Get("/balance", handler.GetBalance)
	api.Get("/transactions", handler.GetTransactions)
	api.Get("/issues", handler.GetIssues)
	
	return handler
}
