package webhooks

import (
	"github.com/neelchoudhary/budgetwallet-api-server/models"
	"github.com/neelchoudhary/budgetwallet-api-server/postgresql"
)

// Service WebhooksService struct
type Service struct {
	txRepo                   postgresql.TxRepository
	financialItemRepo        models.FinancialItemRepository
	financialAccountRepo     models.FinancialAccountRepository
	financialTransactionRepo models.FinancialTransactionRepository
}

// NewWebhooksServer contructor to assign repo
func NewWebhooksServer(txRepo *postgresql.TxRepository, itemRepo *models.FinancialItemRepository, accountRepo *models.FinancialAccountRepository, financialTransactionRepo *models.FinancialTransactionRepository) *Service {
	return &Service{txRepo: *txRepo, financialItemRepo: *itemRepo, financialAccountRepo: *accountRepo, financialTransactionRepo: *financialTransactionRepo}
}
