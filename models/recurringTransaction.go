package models

import (
	"database/sql"
)

// RecurringTransaction ...
type RecurringTransaction struct {
	ID                int64  `json:"id"`
	UserID            int64  `json:"user_id"`
	Name              string `json:"name"`
	CategoryID        int64  `json:"category_id"`
	RecurringCount    int64  `json:"recurring_count"`
	RecurringScore    int64  `json:"recurring_score"`
	IsRecurring       string `json:"is_recurring"`
	RecurringPlaidIDs string `json:"recurring_plaid_ids"` // should be a list []string
}

// RecurringTransactionRepository interface
type RecurringTransactionRepository interface {
	AddRecurringTransaction(tx *sql.Tx, recurringTransaction *RecurringTransaction) error
	GetRecurringTransactions(tx *sql.Tx, userID int64) ([]RecurringTransaction, error)
	RemoveRecurringTransactions(tx *sql.Tx, userID int64) error
}
